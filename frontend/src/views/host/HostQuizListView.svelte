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

    (async function () {
        await loadQuizzes();
    })();
</script>

<div class="p-8 max-w-4xl mx-auto">
    <div class="flex justify-between items-center mb-6">
        <h2 class="text-4xl font-bold">Your quizzes</h2>
        {#if quizzes && quizzes.length > 0}
            <Button on:click={createQuiz}>Create Quiz</Button>
        {/if}
    </div>

    {#if loading}
        <p class="text-gray-500">Loading quizzes...</p>
    {:else if quizzes && quizzes.length > 0}
        <div class="flex flex-col gap-2">
            {#each quizzes as quiz (quiz.id)}
                <QuizCard on:host {quiz} />
            {/each}
        </div>
    {:else}
        <div class="p-8 border-2 border-dashed border-gray-300 rounded-xl text-center bg-gray-50 mt-4">
            <h3 class="text-xl font-semibold text-gray-700 mb-2">No quizzes found</h3>
            <p class="text-gray-500 mb-6">There are currently no quizzes in MongoDB. Create a new quiz to get started!</p>
            <Button on:click={createQuiz}>Create Quiz</Button>
        </div>
    {/if}
</div>

