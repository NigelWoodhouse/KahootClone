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

    $: displayImage = $intermissionImage || $currentQuestion?.image || "";
    $: questionName = $intermissionQuestionName || $currentQuestion?.name || "";
    $: questionAnswer = $currentQuestion?.choices
        .filter((choice) => choice.correct)
        .map((choice) => choice.name)
        .filter((answer) => answer.trim().length > 0)
        .join(", ") || "";

</script>

<div class="page-background relative isolate min-h-screen w-full overflow-hidden flex flex-col text-white">
    <ul class="bubbles" aria-hidden="true">
        {#each [100, 40, 50, 80, 60, 130, 150, 45, 35, 170] as size, i}
            <li
                style="left: {10 + ((i * 37) % 80)}%; width: {size}px; height: {size}px; animation-delay: {i * 1.5}s; animation-duration: {12 + ((i * 7) % 20)}s"
            ></li>
        {/each}
    </ul>
    <!-- Top Header Bar -->
    <div class="relative z-10 flex justify-between items-center px-10 py-6 border-b border-white/10 bg-black/20 backdrop-blur-sm">
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

    <!-- Main Content Stage -->
     <div>
     <div class = "relative z-10 flex flex-col items-center justify-center p-8">

        {#if questionName}
            <div class="rich-text-content text-2xl font-bold tracking-tight text-white/90 drop-shadow">
                {@html sanitizeRichText(questionName)}
            </div>
        {/if}
        {#if questionAnswer}
            <p class="mt-2 text-lg font-semibold text-yellow-200">
                Answer: {questionAnswer}
            </p>
        {/if}
     </div>

    <div class="relative z-10 flex-1 flex items-center justify-center p-8">
        {#if displayImage}
            <div class="w-full max-w-6xl grid grid-cols-1 lg:grid-cols-2 gap-10 items-center">
                <!-- Intermission Image Card -->
                <div class="flex flex-col items-center">
                    <div class="bg-white/10 backdrop-blur-md p-4 rounded-3xl border border-white/20 shadow-2xl overflow-hidden  block flex items-center justify-center transition-all hover:scale-[1.01]">
                        <img
                            src={apiService.getImageUrl(displayImage)}
                            alt="Question intermission visual"
                            class="w-[50vw] h-auto w-auto max-w-[500px] object-contain rounded-2xl shadow-inner"
                        />
                    </div>
                    {#if questionName}
                        <div class="rich-text-content mt-3 text-center text-sm font-medium italic text-white/70">
                            "{@html sanitizeRichText(questionName)}"
                        </div>
                    {/if}
                </div>

                <!-- Leaderboard Section -->
                    <div class="flex flex-col items-center w-full max-w-md rounded-3xl shadow-2xl">
                        <Leaderboard finish={true} leaderboard={$leaderboard} />
                </div>
            </div>
        {:else}
            <!-- If no image, display centered leaderboard -->
                <div class= "flex flex-col items-center w-full max-w-md rounded-3xl shadow-2xl">
                    <Leaderboard finish={true} leaderboard={$leaderboard} />
                </div>
        {/if}
    </div>
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
