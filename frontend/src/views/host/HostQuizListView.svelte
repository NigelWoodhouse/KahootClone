<script lang="ts">
    import Button from "../../lib/Button.svelte";
    import QuizCard from "../../lib/QuizCard.svelte";
    import type { Quiz } from "../../model/quiz";
    import { apiService } from "../../service/api";
    import { push } from "svelte-spa-router";

    let quizzes: Quiz[] = [];
    let loading = true;

    async function loadQuizzes() {
        loading = true;
        quizzes = await apiService.getQuizzes();
        loading = false;
    }

    async function createQuiz() {
        let newQuiz = await apiService.createQuiz();
        if (newQuiz) {
            push(`/edit/${newQuiz.id}`);
        }
    }

    async function onDelete(quiz: Quiz) {
        if (await apiService.deleteQuiz(quiz.id)) {
            quizzes = quizzes.filter((item) => item.id !== quiz.id);
        }
    }

    async function onBack() {
        window.history.back();
    }

    (async function () {
        await loadQuizzes();
    })();
</script>

<div class="lobby-background justify-center min-h-screen max-w-4xl mx-auto">
<div class="flex justify-center items-center gap-4 p-4">
        <h2 class="text-4xl center font-bold text-white">Quizzes</h2>
        {#if quizzes && quizzes.length > 0}
        <div>
            <Button on:click={createQuiz}>Create Quiz</Button>
            <Button on:click={onBack}>Back</Button>
        </div>
        {/if}
</div>


    {#if loading}
        <p class="text-gray-500">Loading quizzes...</p>
    {:else if quizzes && quizzes.length > 0}
        <div class="flex flex-col gap-2">
            {#each quizzes as quiz (quiz.id)}
                <QuizCard on:host on:delete={(event) => onDelete(event.detail)} {quiz} />
            {/each}
        </div>
    {:else}
        <div class="p-8 border-2 border-dashed border-gray-300 rounded-xl text-center bg-gray-50 mt-4">
            <h3 class="text-xl font-semibold text-gray-700 mb-2">No quizzes found</h3>
            <p class="text-gray-500 mb-6">There are currently no quizzes in MongoDB. Create a new quiz to get started!</p>
            <div>
                <Button on:click={createQuiz}>Create Quiz</Button>
                <Button on:click={onBack}>Back</Button>
            </div>
        </div>
    {/if}

</div>

<style>
    .lobby-background {
        background: linear-gradient(to left, #8f94fb, #4e54c8);
    }
</style>