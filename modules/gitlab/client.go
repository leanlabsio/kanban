package gitlab

import (
	"github.com/Unknwon/macaron"
	"github.com/dgrijalva/jwt-go"
	"github.com/pmylund/sortutil"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"net/url"
	"time"
)

type ApiGitlab struct {
	config *Config
	token  []byte
}

type Config struct {
	BasePath string
	Domain   string
	Oauth2   *oauth2.Config
}

// New gitlab api client
func New(c *Config, token string) *ApiGitlab {
	return &ApiGitlab{
		config: c,
		token:  []byte(token),
	}
}

// Get redirect url for gitlab authorisation
func (g *ApiGitlab) OauthUrl(ctx *macaron.Context) {
	ctx.Redirect(g.config.Oauth2.AuthCodeURL("state", oauth2.AccessTypeOffline))
}

// Login with gitlab and get access token
func (g *ApiGitlab) OauthLogin(ctx *macaron.Context, data Oauth2) {
	tok, err := g.config.Oauth2.Exchange(oauth2.NoContext, data.Code)

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["gitlab_token"] = tok.AccessToken
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString(g.token)

	if err != nil {
		log.Printf("%s", err.Error())
	}

	ctx.JSON(http.StatusOK, KbAuth{
		Success: true,
		Token:   tokenString,
	})
}

// Check auth current user and return http client with credential for gitlab request
func (g *ApiGitlab) isAuth(ctx *macaron.Context, log *log.Logger) (*http.Client, error) {
	jwtToken, err := jwt.Parse(ctx.Req.Header["X-Kb-Access-Token"][0], func(token *jwt.Token) (interface{}, error) {
		return []byte(g.token), nil
	})

	if err != nil || !jwtToken.Valid {
		log.Print("%s", err.Error())
		return nil, err
	}

	gitlab_token, _ := jwtToken.Claims["gitlab_token"].(string)
	exp, _ := jwtToken.Claims["exp"].(time.Time)

	tok := &oauth2.Token{
		AccessToken: gitlab_token,
		Expiry:      exp,
	}

	return g.config.Oauth2.Client(oauth2.NoContext, tok), nil
}

// List projects from gitlab
func (g *ApiGitlab) ListProjects(ctx *macaron.Context, log *log.Logger) {
	cl, err := g.isAuth(ctx, log)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, ApiErr{
			Success: false,
			Message: err.Error(),
		})
	}

	path := g.GetUrl([]string{"projects"})

	req, _ := http.NewRequest("GET", path, nil)
	q := req.URL.Query()
	q.Add("per_page", "100")
	q.Add("page", "1")
	req.URL.RawQuery = q.Encode()

	var ret ProjectListResponse
	if err := g.Do(cl, req, &ret.Data); err != nil {
		g.SendError(ctx, err)
	}
	ctx.JSON(http.StatusOK, ret)
}

// Get single project from gitlab
func (g *ApiGitlab) SingleProjects(ctx *macaron.Context, log *log.Logger) {
	cl, err := g.isAuth(ctx, log)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, ApiErr{
			Success: false,
			Message: err.Error(),
		})
	}

	path := g.GetUrl([]string{"projects", url.QueryEscape(ctx.Query("project_id"))})

	req, _ := http.NewRequest("GET", path, nil)

	var ret ProjectSingleResponse
	if err := g.Do(cl, req, &ret.Data); err != nil {
		g.SendError(ctx, err)
	}
	ctx.JSON(http.StatusOK, ret)
}

// Get list issues for gitlab projects
func (g *ApiGitlab) ListIssues(ctx *macaron.Context, log *log.Logger) {
	cl, err := g.isAuth(ctx, log)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, ApiErr{
			Success: false,
			Message: err.Error(),
		})
	}

	path := g.GetUrl([]string{"projects", url.QueryEscape(ctx.Query("project_id")), "issues"})

	req, _ := http.NewRequest("GET", path, nil)
	q := req.URL.Query()
	q.Add("per_page", "200")
	q.Add("page", "1")
	req.URL.RawQuery = q.Encode()

	var ret IssueListResponse
	if err := g.Do(cl, req, &ret.Data); err != nil {
		g.SendError(ctx, err)
	}
	ctx.JSON(http.StatusOK, ret)
}

// Get list milestones for gitlab projects
func (g *ApiGitlab) ListMilestones(ctx *macaron.Context, log *log.Logger) {
	cl, err := g.isAuth(ctx, log)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, ApiErr{
			Success: false,
			Message: err.Error(),
		})
	}

	path := g.GetUrl([]string{"projects", url.QueryEscape(ctx.Query("project_id")), "milestones"})

	req, _ := http.NewRequest("GET", path, nil)
	q := req.URL.Query()
	q.Add("per_page", "100")
	q.Add("page", "1")
	req.URL.RawQuery = q.Encode()

	var ret MilestoneListResponse
	if err := g.Do(cl, req, &ret.Data); err != nil {
		g.SendError(ctx, err)
	}
	ctx.JSON(http.StatusOK, ret)
}

// Get list project members for gitlab projects
func (g *ApiGitlab) ListProjectMembers(ctx *macaron.Context, log *log.Logger) {
	cl, err := g.isAuth(ctx, log)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, ApiErr{
			Success: false,
			Message: err.Error(),
		})
	}

	path := g.GetUrl([]string{"projects", url.QueryEscape(ctx.Query("project_id")), "members"})

	req, _ := http.NewRequest("GET", path, nil)
	q := req.URL.Query()
	q.Add("per_page", "100")
	q.Add("page", "1")
	req.URL.RawQuery = q.Encode()

	var ret MemberListResponse
	if err := g.Do(cl, req, &ret.Data); err != nil {
		g.SendError(ctx, err)
	}
	ctx.JSON(http.StatusOK, ret.Data)
}

// Get list labels for gitlab projects
func (g *ApiGitlab) ListLabels(ctx *macaron.Context, log *log.Logger) {
	cl, err := g.isAuth(ctx, log)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, ApiErr{
			Success: false,
			Message: err.Error(),
		})
	}

	path := g.GetUrl([]string{"projects", url.QueryEscape(ctx.Query("board_id")), "labels"})

	req, _ := http.NewRequest("GET", path, nil)

	var ret LabelListResponse
	if err := g.Do(cl, req, &ret.Data); err != nil {
		g.SendError(ctx, err)
	}
	ctx.JSON(http.StatusOK, ret)
}

// Get list comeent for gitlab issue
func (g *ApiGitlab) ListComments(ctx *macaron.Context, log *log.Logger) {
	cl, err := g.isAuth(ctx, log)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, ApiErr{
			Success: false,
			Message: err.Error(),
		})
	}

	path := g.GetUrl([]string{"projects", url.QueryEscape(ctx.Query("project_id")), "issues", ctx.Query("issue_id"), "notes"})

	req, _ := http.NewRequest("GET", path, nil)

	var ret CommentListResponse
	if err := g.Do(cl, req, &ret.Data); err != nil {
		g.SendError(ctx, err)
	}

	sortutil.AscByField(ret.Data, "CreatedAt")
	ctx.JSON(http.StatusOK, ret)
}
