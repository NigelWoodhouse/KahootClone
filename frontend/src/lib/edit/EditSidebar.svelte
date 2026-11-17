<script lang="ts">
    import { createDefaultQuestion, type QuizQuestion } from "../../model/quiz";
    import Button from "../Button.svelte";
    import SidebarItem from "./SidebarItem.svelte";

    export let questions: QuizQuestion[];
    export let selectedQuestion: QuizQuestion | null;

    function addNew() {
        questions = [...questions, createDefaultQuestion()];
    }
</script>

<div class="page-background min-h-screen w-full min-w-0 overflow-hidden p-2">
    {#each questions as question}
        <SidebarItem {question} bind:selectedQuestion />
    {/each}
    <div class="items-center w-full flex justify-center">
        <Button on:click={addNew}>Add new</Button>
    </div>
</div>
