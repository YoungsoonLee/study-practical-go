package main

import (
	"context"
	"net"
	"testing"
)

func TestRepoService(t *testing.T) {
	s, l := startTestGrpcServer()
	defer s.GracefullStop()

	bufConnDialer := func(ctx context.Context, addr string) (net.Conn, error) {
		return l.Dial()
	}

	client, err := grpc.DialContext(
		context.Background(),
		"", grpc.WithInsecure(),
		grpc.WithContextDialer(bufConnDialer),
	)

	if err != nil {
		t.Fatal(err)
	}

	repoClient := svc.NewRepoClient(client)
	resp, err := respoClient.GetRepos(
		context.Background(),
		&svc.RepoGetRequest{
			CreatedId: "user-123",
			Id:        "repo-123",
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.Repo) != 1 {
		t.Fatalf(
			"Expected to get back 1 repo, got back: %d repos", len(resp.Repo),
		)
	}

	gotID := resp.Repo[0].Id
	gotOwnerID := resp.Repo[0].Owner.Id

	if gotID != "repo-123" {
		t.Errorf("Expected Repo ID to be: repo-123, Got: %s", gotID)
	}
	if gotOwnerID != "user-123" {
		t.Errorf("Expected Creator ID to be: user-123, Got: %s", gotOwnerID)
	}
}
