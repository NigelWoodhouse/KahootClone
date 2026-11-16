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

<div class="page-background min-h-screen px-6 py-6">
    <header class="mx-auto flex w-full max-w-6xl items-center justify-between gap-6">
        <h2 class="text-4xl font-bold text-white">Quizzes</h2>
        <div class="flex items-center justify-end gap-3 rounded-xl">
            <Button on:click={createQuiz}>Create Quiz</Button>
            <Button on:click={onBack}>Back</Button>
        </div>
    </header>

    {#if loading}
        <p class="mx-auto mt-8 max-w-6xl text-white/80">Loading quizzes...</p>
    {:else if quizzes && quizzes.length > 0}
        <div class="mx-auto mt-8 flex w-full max-w-6xl flex-col gap-3">
            {#each quizzes as quiz (quiz.id)}
                <QuizCard on:host on:delete={(event) => onDelete(event.detail)} {quiz} />
            {/each}
        </div>
    {:else}
        <div class="mx-auto mt-8 max-w-6xl rounded-xl border border-dashed border-white/30 bg-white/10 p-8 text-center">
            <h3 class="mb-2 text-xl font-semibold text-white">No quizzes found</h3>
            <p class="text-white/80">There are currently no quizzes in MongoDB. Create a new quiz to get started!</p>
        </div>
    {/if}
</div>
