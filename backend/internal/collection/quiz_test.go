package collection

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"quiz.com/quiz/internal/entity"
)

func TestQuestionImagesUseDistinctBSONFields(t *testing.T) {
	question := entity.QuizQuestion{
		QuestionImage:     "question.png",
		IntermissionImage: "intermission.png",
	}

	encoded, err := bson.Marshal(question)
	if err != nil {
		t.Fatal(err)
	}

	var stored bson.M
	if err := bson.Unmarshal(encoded, &stored); err != nil {
		t.Fatal(err)
	}
	if stored["questionImage"] != "question.png" {
		t.Errorf("questionImage = %v, want question.png", stored["questionImage"])
	}
	if stored["intermissionImage"] != "intermission.png" {
		t.Errorf("intermissionImage = %v, want intermission.png", stored["intermissionImage"])
	}
	if _, exists := stored["image"]; exists {
		t.Error("new image fields must not be encoded using the legacy image key")
	}
}
