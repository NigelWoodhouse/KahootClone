<script lang="ts">
    import type { QuizQuestion } from "../../model/quiz";

    export let question: QuizQuestion;
    export let selectedQuestion: QuizQuestion | null;

    function onClick() {
        selectedQuestion = question;
    }

    $: borderColor =
        selectedQuestion?.id == question.id ? "border-gray-500" : "";
</script>

<button
    class="bg-white rounded-md p-2 border {borderColor} hover:bg-gray-50 flex items-center justify-between text-left transition-colors"
    on:click={onClick}
>
    <span class="truncate">{question.name || "Untitled Question"}</span>
    <!-- {#if question.image}
        <span title="Has intermission image" class="text-purple-600 ml-1.5 flex-shrink-0">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
            </svg>
        </span>
    {/if} -->
</button>

