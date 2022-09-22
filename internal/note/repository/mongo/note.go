package mongo

import (
	"context"

	"github.com/lekan-rust/notes-app/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Note struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      primitive.ObjectID `bson:"userId"`
	Title       string             `bson:"title"`
	Description string             `bson"description"`
}

func toMongoNote(n *models.Note) *Note {
	uid, _ := primitive.ObjectIDFromHex(n.UserID)

	return &Note{
		UserID:      uid,
		Title:       n.Title,
		Description: n.Description,
	}
}

func toModelNote(n *Note) *models.Note {
	return &models.Note{
		ID:          n.ID.Hex(),
		UserID:      n.UserID.Hex(),
		Title:       n.Title,
		Description: n.Description,
	}
}

func toModelNotes(notes []*Note) []*models.Note {
	out := make([]*models.Note, len(notes))

	for i, n := range notes {
		out[i] = toModelNote(n)
	}
	return out
}

type NoteRepository struct {
	db *mongo.Collection
}

func NewNoteRepository(db *mongo.Database, collection string) *NoteRepository {
	return &NoteRepository{
		db: db.Collection(collection),
	}
}

// CreateNote is a function to create a new note in database
func (r NoteRepository) CreateNote(ctx context.Context, user *models.User, note *models.Note) error {
	note.UserID = user.ID

	model := toMongoNote(note)

	res, err := r.db.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	note.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return nil
}

// GetNotes returns all notes of some user
func (r NoteRepository) GetNotes(ctx context.Context, user *models.User) ([]*models.Note, error) {
	uid, _ := primitive.ObjectIDFromHex(user.ID)
	cur, err := r.db.Find(ctx, bson.M{
		"userId": uid,
	})
	defer cur.Close(ctx)

	if err != nil {
		return nil, err
	}

	out := make([]*Note, 0)

	for cur.Next(ctx) {
		user := new(Note)
		err := cur.Decode(user)
		if err != nil {
			return nil, err
		}

		out = append(out, user)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return toModelNotes(out), nil
}

// DeleteNote removes one note of some user
func (r NoteRepository) DeleteNote(ctx context.Context, user *models.User, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	uID, _ := primitive.ObjectIDFromHex(user.ID)

	_, err := r.db.DeleteOne(ctx, bson.M{"_id": objID, "iserId": uID})
	return err
}
