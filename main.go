package main

import (
	"context"
	"entedge/ent"
	"entedge/ent/group"
	"entedge/ent/image"
	"entedge/ent/pet"
	"entedge/ent/user"
	"entedge/infrastructure"
	"io/ioutil"
	"log"
)

func main() {
	conn := infrastructure.OpenDB()
	defer conn.Close()

	// Run the auto migration tool.
	if err := conn.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// GetImage(context.Background(), conn)
	// InsertImage(context.Background(), conn)
	Do(context.Background(), conn)
	// GetGroupUser(context.Background(), conn)
}

func GetGroupUser(ctx context.Context, client *ent.Client) {
	// groupのIDが1001のuserとgroupのnameを取得する
	users, err := client.User.Query().
		Where(
			user.HasGroupsWith(
				group.ID(11),
			),
		).
		WithGroups().
		All(ctx)
	if err != nil {
		log.Fatalf("failed querying user names: %v", err)
	}

	log.Printf("users: %+v", users)
	for _, u := range users {
		log.Printf("User ID: %d, Name: %s, Group Name: %s\n", u.ID, u.Name, u.Edges.Groups[0].Name)
	}
}

func Do(ctx context.Context, client *ent.Client) {
	// ownerのIDが1001のpetとuserのnameを取得する
	pets, err := client.Pet.Query().
		Where(
			pet.HasOwnerWith(
				user.ID(1001),
			),
		).
		WithOwner().
		All(ctx)
	if err != nil {
		log.Fatalf("failed querying pet names: %v", err)
	}
	log.Printf("pets: %+v", pets)
	for _, p := range pets {
		log.Printf("Pet ID: %d, Name: %s, Owner ID: %d, Owner Name: %s\n", p.ID, p.Name, p.Edges.Owner.ID, p.Edges.Owner.Name)
	}
}

func Join(ctx context.Context, client *ent.Client) {
	//
}

// imageテーブルにtest.pngを挿入する
func InsertImage(ctx context.Context, client *ent.Client) {
	imageFilePath := "test.png"
	imageBytes, err := ioutil.ReadFile(imageFilePath)
	if err != nil {
		log.Fatalf("failed reading image file: %v", err)
	}

	_, err = client.Image.Create().
		SetData(imageBytes).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed inserting image: %v", err)
	}

	log.Printf("inserted image: %s", imageFilePath)
}

// バイナリデータを取得し、image.pngとして保存する
func GetImage(ctx context.Context, client *ent.Client) {
	image, err := client.Image.Query().
		Where(image.ID(1)).
		Only(ctx)
	if err != nil {
		log.Fatalf("failed querying image: %v", err)
	}

	err = ioutil.WriteFile("image.png", image.Data, 0644)
	if err != nil {
		log.Fatalf("failed writing image file: %v", err)
	}

	log.Printf("wrote image: image.png")
}
