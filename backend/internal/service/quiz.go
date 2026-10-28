package service

import (
	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"quiz.com/quiz/internal/collection"
	"quiz.com/quiz/internal/entity"
)

type QuizService struct {
	quizCollection *collection.QuizCollection
}

func Quiz(quizCollection *collection.QuizCollection) *QuizService {
	return &QuizService{
		quizCollection: quizCollection,
	}
}

func (s QuizService) NewDefaultQuestion() entity.QuizQuestion {
	return entity.QuizQuestion{
		Id:    uuid.New().String(),
		Name:  "Question",
		Time:  60,
		Image: "",
		Choices: []entity.QuizChoice{
			{
				Id:      uuid.New().String(),
				Name:    "",
				Correct: false,
			},
			{
				Id:      uuid.New().String(),
				Name:    "",
				Correct: false,
			},
			{
				Id:      uuid.New().String(),
				Name:    "",
				Correct: false,
			},
			{
				Id:      uuid.New().String(),
				Name:    "",
				Correct: false,
			},
		},
	}
}

func (s QuizService) CreateQuiz(name string, questions []entity.QuizQuestion) (*entity.Quiz, error) {
	if name == "" {
		name = "New Quiz"
	}

	if len(questions) == 0 {
		questions = []entity.QuizQuestion{
			s.NewDefaultQuestion(),
		}
	}

	quiz := entity.Quiz{
		Id:        primitive.NewObjectID(),
		Name:      name,
		Questions: questions,
	}

	return s.quizCollection.InsertQuiz(quiz)
}

func (s QuizService) GetQuizById(id primitive.ObjectID) (*entity.Quiz, error) {
	return s.quizCollection.GetQuizById(id)
}

func (s QuizService) UpdateQuiz(id primitive.ObjectID, name string, questions []entity.QuizQuestion) error {
	quiz, err := s.quizCollection.GetQuizById(id)
	if err != nil {
		return err
	}

	if quiz == nil {
		return errors.New("quiz not found")
	}

	quiz.Name = name
	quiz.Questions = questions
	return s.quizCollection.UpdateQuiz(*quiz)
}

func (s QuizService) DeleteQuiz(id primitive.ObjectID) (bool, error) {
	return s.quizCollection.DeleteQuiz(id)
}

func (s QuizService) GetQuizzes() ([]entity.Quiz, error) {
	return s.quizCollection.GetQuizzes()
}
