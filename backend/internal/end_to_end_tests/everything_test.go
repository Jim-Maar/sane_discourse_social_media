package end_to_end_tests

import (
	"encoding/json"
	"net/http"
	"sane-discourse-backend/internal/handlers"
	"sane-discourse-backend/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEverything(t *testing.T) {
	r := SetupTestServer()

	// Test User Login
	userName := "Tim"
	userEmail := "Tim@Tom.com"
	loginRequest := handlers.LoginUserRequest{
		Name:  userName,
		Email: userEmail,
	}
	userResponse := PerformRequest(r, "PUT", "/auth/login", loginRequest)
	user := models.User{}
	err := json.Unmarshal(userResponse.Body.Bytes(), &user)
	if err != nil {
		t.Fatalf("Failed to unmarshal user response: %v", err)
	}
	assert.Equal(t, http.StatusOK, userResponse.Result().StatusCode)
	assert.Equal(t, userName, user.Username)
	assert.Equal(t, userEmail, user.Email)
	assert.NotNil(t, user.ID)

	// Test Create Post from URL
	postURL := "https://forum.effectivealtruism.org/posts/hkimyETEo76hJ6NpW/on-caring"
	createPostRequest := handlers.CreatePostRequest{
		URL: postURL,
	}
	createPostResponse := PerformRequest(r, "PUT", "/user/posts/create", createPostRequest)
	post := models.Post{}
	err = json.Unmarshal(createPostResponse.Body.Bytes(), &post)
	if err != nil {
		t.Fatalf("Failed to unmarshal create post response: %v", err)
	}
	assert.Equal(t, "On Caring — EA Forum", post.Title)
	assert.Equal(t, postURL, post.URL)
	assert.Equal(t, "article", post.Type)
	assert.Equal(t, "", post.Author)
	assert.NotNil(t, post.ID)
	assert.Equal(t, http.StatusOK, createPostResponse.Result().StatusCode)

	// Test Add Post
	addPostRequest := handlers.AddPostRequest{
		Post: post,
	}
	addPostResponse := PerformRequest(r, "PUT", "/user/posts/add", addPostRequest)
	addedPost := models.Post{}
	err = json.Unmarshal(addPostResponse.Body.Bytes(), &addedPost)
	if err != nil {
		t.Fatalf("Failed to unmarshal add post response: %v", err)
	}
	assert.Equal(t, post.Title, addedPost.Title)
	assert.Equal(t, post.URL, addedPost.URL)
	assert.Equal(t, post.Type, addedPost.Type)
	assert.Equal(t, post.Author, addedPost.Author)
	assert.NotNil(t, addedPost.ID)
	assert.Equal(t, http.StatusOK, addPostResponse.Result().StatusCode)

	// Test Get User Posts
	getUserPostsResponse := PerformRequest(r, "GET", "/user/posts", nil)
	var userPosts []models.Post
	err = json.Unmarshal(getUserPostsResponse.Body.Bytes(), &userPosts)
	if err != nil {
		t.Fatalf("Failed to unmarshal user posts response: %v", err)
	}
	assert.Equal(t, http.StatusOK, getUserPostsResponse.Result().StatusCode)
	assert.Equal(t, len(userPosts), 1, "User should have 1 post")

	// Test Get Feed
	getFeedResponse := PerformRequest(r, "GET", "/home", nil)
	var feedPosts []models.Post
	err = json.Unmarshal(getFeedResponse.Body.Bytes(), &feedPosts)
	if err != nil {
		t.Fatalf("Failed to unmarshal feed response: %v", err)
	}
	assert.Equal(t, http.StatusOK, getFeedResponse.Result().StatusCode)
	assert.GreaterOrEqual(t, len(feedPosts), 1, "Feed should have at least 1 post")

	// Test Add Header Component to Userpage
	addHeaderRequest := handlers.AddComponentRequest{
		Index: 0,
		Component: models.Component{
			Header: &models.HeaderComponent{
				Content: "Welcome to My Page",
				Size:    models.HeaderComponentSizeLarge,
			},
		},
	}
	addHeaderResponse := PerformRequest(r, "PUT", "/userpage/component/add", addHeaderRequest)
	var userpageAfterHeader models.Userpage
	err = json.Unmarshal(addHeaderResponse.Body.Bytes(), &userpageAfterHeader)
	if err != nil {
		t.Fatalf("Failed to unmarshal userpage response: %v", err)
	}
	assert.Equal(t, http.StatusOK, addHeaderResponse.Result().StatusCode)
	assert.Equal(t, user.ID, userpageAfterHeader.UserID)
	assert.NotNil(t, userpageAfterHeader.Components[0].Header)

	// Test Add Paragraph Component to Userpage
	addParagraphRequest := handlers.AddComponentRequest{
		Index: 1,
		Component: models.Component{
			Paragraph: &models.PragraphComponent{
				Content: "This is a paragraph describing my interests.",
			},
		},
	}
	addParagraphResponse := PerformRequest(r, "PUT", "/userpage/component/add", addParagraphRequest)
	var userpageAfterParagraph models.Userpage
	err = json.Unmarshal(addParagraphResponse.Body.Bytes(), &userpageAfterParagraph)
	if err != nil {
		t.Fatalf("Failed to unmarshal userpage response: %v", err)
	}
	assert.Equal(t, http.StatusOK, addParagraphResponse.Result().StatusCode)
	assert.NotNil(t, userpageAfterParagraph.Components[1].Paragraph)

	// Test Add Post Component to Userpage
	addPostComponentRequest := handlers.AddComponentRequest{
		Index: 2,
		Component: models.Component{
			Post: &models.PostComponent{
				PostID: addedPost.ID,
				Size:   models.PostComponentSizeLarge,
			},
		},
	}
	addPostComponentResponse := PerformRequest(r, "PUT", "/userpage/component/add", addPostComponentRequest)
	var userpageAfterPostComponent models.Userpage
	err = json.Unmarshal(addPostComponentResponse.Body.Bytes(), &userpageAfterPostComponent)
	if err != nil {
		t.Fatalf("Failed to unmarshal userpage response: %v", err)
	}
	assert.Equal(t, http.StatusOK, addPostComponentResponse.Result().StatusCode)
	assert.NotNil(t, userpageAfterPostComponent.Components[2].Post)

	// Test Add Divider Component to Userpage
	addDividerRequest := handlers.AddComponentRequest{
		Index: 3,
		Component: models.Component{
			Divider: &models.DividerComponent{
				Style: models.RegularDevider,
			},
		},
	}
	addDividerResponse := PerformRequest(r, "PUT", "/userpage/component/add", addDividerRequest)
	var userpageAfterDivider models.Userpage
	err = json.Unmarshal(addDividerResponse.Body.Bytes(), &userpageAfterDivider)
	if err != nil {
		t.Fatalf("Failed to unmarshal userpage response: %v", err)
	}
	assert.Equal(t, http.StatusOK, addDividerResponse.Result().StatusCode)
	assert.NotNil(t, userpageAfterDivider.Components[3].Divider)

	// Test Move Component in Userpage
	moveComponentRequest := handlers.MoveComponentRequest{
		PrevIndex: 3,
		NewIndex:  1,
	}
	moveComponentResponse := PerformRequest(r, "PUT", "/userpage/component/move", moveComponentRequest)
	var userpageAfterMove models.Userpage
	err = json.Unmarshal(moveComponentResponse.Body.Bytes(), &userpageAfterMove)
	if err != nil {
		t.Fatalf("Failed to unmarshal userpage response: %v", err)
	}
	assert.Equal(t, http.StatusOK, moveComponentResponse.Result().StatusCode)
	assert.NotNil(t, userpageAfterMove.Components[1].Divider)
	assert.NotNil(t, userpageAfterMove.Components[2].Paragraph)
	assert.NotNil(t, userpageAfterMove.Components[3].Post)
}
