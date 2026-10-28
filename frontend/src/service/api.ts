import { createDefaultQuestion, type Quiz } from "../model/quiz";

export class ApiService {
    async getQuizById(id: string): Promise<Quiz | null> {
        let response = await fetch(`http://localhost:3000/api/quizzes/${id}`);
        if (!response.ok) {
            return null;
        }

        let json = await response.json();
        return json;
    }

    async getQuizzes(): Promise<Quiz[]> {
        let response = await fetch("http://localhost:3000/api/quizzes");
        if (!response.ok) {
            alert("Failed to fetch quizzes!");
            return [];
        }

        let json = await response.json();
        return json || [];
    }

    async createQuiz(name: string = "New Quiz"): Promise<Quiz | null> {
        let response = await fetch("http://localhost:3000/api/quizzes", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                name,
                questions: [createDefaultQuestion()]
            })
        });

        if (!response.ok) {
            alert("Failed to create quiz!");
            return null;
        }

        let json = await response.json();
        return json;
    }

    async saveQuiz(quizId: string, quiz: Quiz) {
        let response = await fetch(`http://localhost:3000/api/quizzes/${quizId}`, {
            method: "PUT",
            body: JSON.stringify(quiz),
            headers: {
                "Content-Type": "application/json"
            }
        });

        if (!response.ok) {
            alert("Failed to save quiz!");
            return;
        }
    }

    async deleteQuiz(quizId: string): Promise<boolean> {
        try {
            const response = await fetch(`http://localhost:3000/api/quizzes/${quizId}`, {
                method: "DELETE"
            });

            if (!response.ok) {
                alert("Failed to delete quiz!");
                return false;
            }

            return true;
        } catch (error) {
            alert(`Failed to delete quiz: ${error instanceof Error ? error.message : error}`);
            return false;
        }
    }

    async uploadImage(file: File): Promise<{ id: string; url: string } | null> {
        const formData = new FormData();
        formData.append("image", file);

        try {
            const response = await fetch("http://localhost:3000/api/upload", {
                method: "POST",
                body: formData
            });

            if (!response.ok) {
                const errText = await response.text();
                alert(`Failed to upload image: ${errText}`);
                return null;
            }

            return await response.json();
        } catch (err: any) {
            alert(`Error uploading image: ${err?.message || err}`);
            return null;
        }
    }

    getImageUrl(imagePath?: string | null): string {
        if (!imagePath) return "";
        if (
            imagePath.startsWith("http://") ||
            imagePath.startsWith("https://") ||
            imagePath.startsWith("data:")
        ) {
            return imagePath;
        }
        if (imagePath.startsWith("/")) {
            return `http://localhost:3000${imagePath}`;
        }
        return `http://localhost:3000/api/images/${imagePath}`;
    }
}

export const apiService = new ApiService();
