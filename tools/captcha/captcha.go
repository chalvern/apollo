package captcha

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/chalvern/sugar"

	"github.com/chalvern/apollo/tools/rand"

	"github.com/gin-contrib/cache/persistence"
)

var (
	defaultChars = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
)

const (
	// default captcha attributes
	challengeNums    = 6
	expiration       = 600 * time.Second
	fieldIDName      = "captcha_id"
	fieldCaptchaName = "captcha"
	cachePrefix      = "captcha_"
	defaultURLPrefix = "/captcha/"
)

// Captcha struct
type Captcha struct {
	// beego cache store
	store persistence.CacheStore

	// url prefix for captcha image
	URLPrefix string

	// specify captcha id input field name
	FieldIDName string
	// specify captcha result input field name
	FieldCaptchaName string

	// captcha image width and height
	StdWidth  int
	StdHeight int

	// captcha chars nums
	ChallengeNums int

	// captcha expiration seconds
	Expiration time.Duration

	// cache key prefix
	CachePrefix string
}

// generate key string
func (c *Captcha) key(id string) string {
	return c.CachePrefix + id
}

// generate rand chars with default chars
func (c *Captcha) genRandChars() []byte {
	return rand.RandomCreateBytes(c.ChallengeNums, defaultChars...)
}

// CreateCaptchaHTML template func for output html
func (c *Captcha) CreateCaptchaHTML() template.HTML {
	value, err := c.CreateCaptcha()
	if err != nil {
		sugar.Errorf("Create Captcha Error:", err)
		return ""
	}

	// create html
	return template.HTML(fmt.Sprintf(`<input type="hidden" name="%s" value="%s">`+
		`<a class="captcha" href="javascript:">`+
		`<img onclick="this.src=('%s%s.png?reload='+(new Date()).getTime())" class="captcha-img" src="%s%s.png">`+
		`</a>`, c.FieldIDName, value, c.URLPrefix, value, c.URLPrefix, value))
}

// CreateCaptcha create a new captcha id
func (c *Captcha) CreateCaptcha() (string, error) {
	// generate captcha id
	id := string(rand.RandomCreateBytes(15))

	// get the captcha chars
	chars := c.genRandChars()

	// save to store
	if err := c.store.Set(c.key(id), chars, c.Expiration); err != nil {
		return "", err
	}

	return id, nil
}

// VerifyReq verify from a request
func (c *Captcha) VerifyReq(req *http.Request) bool {
	req.ParseForm()
	return c.Verify(req.Form.Get(c.FieldIDName), req.Form.Get(c.FieldCaptchaName))
}

// Verify direct verify id and challenge string
func (c *Captcha) Verify(id string, challenge string) (success bool) {
	if len(challenge) == 0 || len(id) == 0 {
		return
	}

	var chars []byte

	key := c.key(id)

	value := []byte{}
	if err := c.store.Get(key, &value); err == nil {
		chars = value
	} else {
		return
	}

	defer func() {
		// finally remove it
		c.store.Delete(key)
	}()

	if len(chars) != len(challenge) {
		return
	}
	// verify challenge
	for i, c := range chars {
		if c != challenge[i]-48 {
			return
		}
	}

	return true
}

// NewCaptcha create a new captcha.Captcha
func NewCaptcha(urlPrefix string, store persistence.CacheStore) *Captcha {
	cpt := &Captcha{}
	cpt.store = store
	cpt.FieldIDName = fieldIDName
	cpt.FieldCaptchaName = fieldCaptchaName
	cpt.ChallengeNums = challengeNums
	cpt.Expiration = expiration
	cpt.CachePrefix = cachePrefix
	cpt.StdWidth = stdWidth
	cpt.StdHeight = stdHeight

	if len(urlPrefix) == 0 {
		urlPrefix = defaultURLPrefix
	}

	if urlPrefix[len(urlPrefix)-1] != '/' {
		urlPrefix += "/"
	}

	cpt.URLPrefix = urlPrefix

	return cpt
}

// Handler beego filter handler for serve captcha image
func (c *Captcha) Handler(ctx *gin.Context) {
	var chars []byte

	id := path.Base(ctx.Request.RequestURI)
	if i := strings.Index(id, "."); i != -1 {
		id = id[:i]
	}

	key := c.key(id)

	if len(ctx.Query("reload")) > 0 {
		chars = c.genRandChars()
		if err := c.store.Set(key, chars, c.Expiration); err != nil {
			ctx.String(500, "captcha reload error")
			sugar.Errorf("Reload Create Captcha Error:", err)
			return
		}
	} else {
		value := []byte{}
		if err := c.store.Get(key, &value); err == nil {
			chars = value
		} else {
			ctx.String(404, "captcha not found")
			return
		}
	}

	img := NewImage(chars, c.StdWidth, c.StdHeight)
	if _, err := img.WriteTo(ctx.Writer); err != nil {
		sugar.Errorf("Write Captcha Image Error:", err)
	}
}
