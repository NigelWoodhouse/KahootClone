<script lang="ts">
    import QuizChoiceCard from "../../lib/play/QuizChoiceCard.svelte";
    import { COLORS, type QuizChoice } from "../../model/quiz";
    import { tick, currentQuestion, gameState } from "../../service/host/host";
    import { GameState } from "../../service/net";

    function getCardColor(choice: QuizChoice, currentGameState: GameState, defaultColor: string){
        if(currentGameState != GameState.Reveal)
            return defaultColor;

        return choice.correct ? "bg-green-400" : "bg-red-400";
    }

    // export let game: HostGame;
</script>

<div class="page-background relative isolate min-h-screen w-full overflow-hidden flex flex-col text-white">
    <ul class="bubbles" aria-hidden="true">
        {#each [100, 40, 50, 80, 60, 130, 150, 45, 35, 170] as size, i}
            <li
                style="left: {10 + ((i * 37) % 80)}%; width: {size}px; height: {size}px; animation-delay: {i * 1.5}s; animation-duration: {12 + ((i * 7) % 20)}s"
            ></li>
        {/each}
    </ul>
{#if $currentQuestion != null}
        {#if $tick > 0}
            <div class="absolute top-6 right-6 z-10 flex items-center gap-2 rounded-full border border-white/15 bg-white/10 px-4 py-2 text-sm font-semibold">
                <svg class="w-6 h-6 text-yellow-300 animate-spin-slow" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <circle cx="12" cy="12" r="10" stroke-width="2" />
                    <polyline points="12 6 12 12 16 14" stroke-width="2" stroke-linecap="round" />
                </svg>
                <span class="text-2xl font-bold">{$tick}s</span>
            </div>
        {/if}
        <div class="relative z-10 flex flex-1 items-center justify-center px-6 text-center">
            <h1 class="text-6xl font-bold tracking-tight text-white drop-shadow">
                {$currentQuestion.name}
            </h1>
        </div>
        <div class="flex flex-wrap w-full h-96">
            {#each COLORS as color, i}
                <QuizChoiceCard color={getCardColor($currentQuestion.choices[i], $gameState, color)}>
                    <p class="pl-14">{$currentQuestion.choices[i].name}</p>
                </QuizChoiceCard>
            {/each}
        </div>
{/if}
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