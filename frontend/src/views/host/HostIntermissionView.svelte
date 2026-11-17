<script lang="ts">
    import Button from "../../lib/Button.svelte";
    import Leaderboard from "../../lib/Leaderboard.svelte";
    import { HostGame, intermissionImage, leaderboard, intermissionQuestionName, currentQuestion, tick } from "../../service/host/host";
    import { apiService } from "../../service/api";
    import { sanitizeRichText } from "../../lib/richText";

    export let game: HostGame;

    function skip() {
        game.start();
    }

    $: displayImage = $intermissionImage || $currentQuestion?.intermissionImage || "";
    $: questionName = $intermissionQuestionName || $currentQuestion?.name || "";
    $: questionAnswer = $currentQuestion?.choices
        .filter((choice) => choice.correct)
        .map((choice) => choice.name)
        .filter((answer) => answer.trim().length > 0)
        .join(", ") || "";

</script>

<div class="page-background relative isolate flex h-dvh min-h-0 w-full flex-col overflow-hidden text-white">
    <ul class="bubbles" aria-hidden="true">
        {#each [100, 40, 50, 80, 60, 130, 150, 45, 35, 170] as size, i}
            <li
                style="left: {10 + ((i * 37) % 80)}%; width: {size}px; height: {size}px; animation-delay: {i * 1.5}s; animation-duration: {12 + ((i * 7) % 20)}s"
            ></li>
        {/each}
    </ul>
    <!-- Top Header Bar -->
    <div class="relative z-10 flex shrink-0 items-center justify-between border-b border-white/10 bg-black/20 px-10 py-6 backdrop-blur-sm">
        <div class="flex items-center gap-4">
            <span class="bg-yellow-400 text-purple-950 px-3.5 py-1 rounded-full text-sm font-extrabold tracking-wider uppercase shadow-md">
                Intermission
            </span>
        </div>
        <div class="flex items-center gap-4">
            {#if $tick > 0}
                <div class="flex items-center gap-1.5 bg-white/10 px-4 py-1.5 rounded-full border border-white/15 text-sm font-semibold">
                    <svg class="w-4 h-4 text-yellow-300 animate-spin-slow" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <circle cx="12" cy="12" r="10" stroke-width="2" />
                        <polyline points="12 6 12 12 16 14" stroke-width="2" stroke-linecap="round" />
                    </svg>
                    <span>{$tick}s</span>
                </div>
            {/if}
            <Button on:click={skip}>Next Question</Button>
        </div>
    </div>

    <div class="relative z-10 flex h-[20vh] min-h-24 max-h-48 shrink-0 flex-col items-center justify-center overflow-y-auto px-6 py-3 text-center">
        {#if questionName}
            <div class="rich-text-content text-2xl font-bold tracking-tight text-white/90 drop-shadow">
                {@html sanitizeRichText(questionName)}
            </div>
        {/if}
        {#if questionAnswer}
            <p class="mt-2 text-2xl font-semibold text-yellow-200 p-4">
                Answer: {questionAnswer}
            </p>
        {/if}
    </div>

    <div class="relative z-10 flex min-h-0 flex-1 items-center justify-center p-6">
        {#if displayImage}
            <div class="flex h-full min-h-0 w-full max-w-7xl items-center justify-center gap-10">
                <div class="flex h-full min-h-0 min-w-0 flex-1 items-center justify-center">
                    <img
                        src={apiService.getImageUrl(displayImage)}
                        alt="Intermission visual"
                        class="max-h-full max-w-full rounded-2xl border border-white/20 object-contain shadow-2xl"
                    />
                </div>
                <div class="flex min-h-0 w-full max-w-md shrink-0 items-center justify-center">
                    <Leaderboard finish={true} leaderboard={$leaderboard} />
                </div>
            </div>
        {:else}
            <div class="flex w-full max-w-md items-center justify-center">
                <Leaderboard finish={true} leaderboard={$leaderboard} />
            </div>
        {/if}
    </div>

</div>

<style>
    @keyframes spin-slow {
        from { transform: rotate(0deg); }
        to { transform: rotate(360deg); }
    }
    .animate-spin-slow {
        animation: spin-slow 8s linear infinite;
    }
</style>
