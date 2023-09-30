package interfaces

import (
	"context"
	"idid-api/ent"

	"github.com/google/uuid"
)

// PostUseCase is a use case for the Post entity
type PostUseCase struct {
	client *ent.Client
}

// NewPostUseCase creates a new PostUseCase
func NewPostUseCase(client *ent.Client) *PostUseCase {
	return &PostUseCase{client: client}
}

// CreatePost creates a new Post
func (uc *PostUseCase) CreatePost(ctx context.Context, title, content string) *ent.Post {
	post := uc.client.Post.Create().SetTitle(title).SetContent(content).SaveX(ctx)
	return post
}

// GetPost gets a Post by ID
func (uc *PostUseCase) GetPost(ctx context.Context, id uuid.UUID) *ent.Post {
	post, _ := uc.client.Post.Get(ctx, id)
	return post
}

// GetPosts gets all Posts
func (uc *PostUseCase) GetPosts(ctx context.Context) []*ent.Post {
	posts, _ := uc.client.Post.Query().All(ctx)
	return posts
}

// UpdatePost updates a Post by ID
func (uc *PostUseCase) UpdatePost(ctx context.Context, id uuid.UUID, title, content string) *ent.Post {
	post := uc.client.Post.UpdateOneID(id).SetTitle(title).SetContent(content).SaveX(ctx)
	return post
}

// DeletePost deletes a Post by ID
func (uc *PostUseCase) DeletePost(ctx context.Context, id uuid.UUID) {
	uc.client.Post.DeleteOneID(id).ExecX(ctx)
}
