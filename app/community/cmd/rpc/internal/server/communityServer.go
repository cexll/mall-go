package server

import (
	"context"

	"mall-go/app/community/cmd/pb"
	"mall-go/app/community/cmd/rpc/internal/svc"
)

type CommunityServer struct {
	svcCtx *svc.ServiceContext
}

func NewCommunityServer(svcCtx *svc.ServiceContext) *CommunityServer {
	return &CommunityServer{
		svcCtx: svcCtx,
	}
}

func (s *CommunityServer) GetPostList(ctx context.Context, request *pb.GetPostListRequest) (*pb.GetPostListResponse, error) {
	//TODO implement me
	return &pb.GetPostListResponse{}, nil

}

func (s *CommunityServer) GetPostDetail(ctx context.Context, request *pb.GetPostDetailRequest) (*pb.GetPostDetailResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CommunityServer) PushPost(ctx context.Context, request *pb.PushPostRequest) (*pb.PushPostResponse, error) {
	//TODO implement me
	return &pb.PushPostResponse{}, nil
}

func (s *CommunityServer) SetPost(ctx context.Context, request *pb.SetPostRequest) (*pb.SetPostResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CommunityServer) DeletePost(ctx context.Context, request *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CommunityServer) GetCommentList(ctx context.Context, request *pb.GetCommentListRequest) (*pb.GetCommentListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CommunityServer) AddComment(ctx context.Context, request *pb.AddCommentRequest) (*pb.AddCommentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CommunityServer) ReplyComment(ctx context.Context, request *pb.ReplyCommentRequest) (*pb.ReplyCommentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CommunityServer) DeleteComment(ctx context.Context, request *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CommunityServer) CollectPost(ctx context.Context, request *pb.CollectPostRequest) (*pb.CollectPostResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CommunityServer) GetCollectPostList(ctx context.Context, request *pb.GetCollectPostListRequest) (*pb.GetCollectPostListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CommunityServer) StarPost(ctx context.Context, request *pb.StarPostRequest) (*pb.StarPostResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CommunityServer) GetStarPostList(ctx context.Context, request *pb.GetStarPostListRequest) (*pb.GetStarPostListResponse, error) {
	//TODO implement me
	panic("implement me")
}
