<script lang="ts">
    import Button from "../../lib/Button.svelte";
    import EditQuestion from "../../lib/edit/EditQuestion.svelte";
    import EditSidebar from "../../lib/edit/EditSidebar.svelte";
    import type { Quiz, QuizQuestion } from "../../model/quiz";
    import { apiService } from "../../service/api";

    export let params: Record<string, string>;

    let quiz: Quiz | null;
    let selectedQuestion: QuizQuestion | null = null;

    function onQuestionDelete() {
        if (quiz == null) return;
        quiz.questions = quiz.questions.filter(
            (q) => q.id != selectedQuestion?.id,
        );

        selectedQuestion = null;
    }

    function goBack() {
        window.history.back();
    }

    (async function () {
        quiz = await apiService.getQuizById(params["quizId"]);
        if (quiz && quiz.questions && quiz.questions.length > 0) {
            selectedQuestion = quiz.questions[0];
        }
    })();
    
    let isSaving = false;
    let saveSuccess = false;

    async function save() {
        if (quiz == null) return;
        isSaving = true;
        try {
            await apiService.saveQuiz(quiz.id, quiz);
            saveSuccess = true;
            setTimeout(() => {
                saveSuccess = false;
            }, 2500);
        } finally {
            isSaving = false;
        }
    }
</script>

{#if quiz != null}
    <div class="bg-gray-100 w-full p-2 flex justify-end items-center border-b">
        <div class="flex items-center gap-2">
            {#if saveSuccess}
                <span class="text-xs text-emerald-600 font-semibold flex items-center gap-1">
                    <svg class="w-4 h-4 inline" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" /></svg>
                    Saved to MongoDB!
                </span>
            {/if}
            <input
                type="text"
                class="border rounded px-2 py-1 text-sm bg-white"
                placeholder="Quiz name"
                bind:value={quiz.name}
            />
            <Button on:click={save}>{isSaving ? "Saving..." : "Save"}</Button>
            <Button on:click={goBack}>{"Back"}</Button>
        </div>
    </div>
    <div class="flex">
        <EditSidebar bind:questions={quiz.questions} bind:selectedQuestion />
        {#if selectedQuestion != null}
            <EditQuestion
                on:delete={onQuestionDelete}
                on:change={() => quiz = quiz}
                bind:selectedQuestion
            />
        {/if}
    </div>
{:else}
    Quiz not found.
{/if}
