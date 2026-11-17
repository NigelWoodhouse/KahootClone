<script lang="ts">
    import QuizChoiceCard from "../../lib/play/QuizChoiceCard.svelte";
    import { COLORS, type QuizChoice } from "../../model/quiz";
    import { tick, currentQuestion, gameState, questionImage } from "../../service/host/host";
    import { GameState } from "../../service/net";
    import { apiService } from "../../service/api";
    import { sanitizeRichText } from "../../lib/richText";

    function getCardColor(choice: QuizChoice, currentGameState: GameState, defaultColor: string){
        if(currentGameState != GameState.Reveal)
            return defaultColor;

        return choice.correct ? "bg-green-400" : "bg-red-400";
    }

    function fitQuestionText(node: HTMLDivElement) {
        const container = node.parentElement;
        if (!container) return;

        let frame = 0;
        const fit = () => {
            cancelAnimationFrame(frame);
            frame = requestAnimationFrame(() => {
                let min = 14;
                let max = 60;
                let best = min;

                while (min <= max) {
                    const size = Math.floor((min + max) / 2);
                    node.style.fontSize = `${size}px`;

                    if (node.scrollHeight <= container.clientHeight
                        && node.scrollWidth <= container.clientWidth) {
                        best = size;
                        min = size + 1;
                    } else {
                        max = size - 1;
                    }
                }

                node.style.fontSize = `${best}px`;
            });
        };

        const resizeObserver = new ResizeObserver(fit);
        resizeObserver.observe(container);
        const mutationObserver = new MutationObserver(fit);
        mutationObserver.observe(node, { childList: true, characterData: true, subtree: true });
        fit();

        return {
            destroy() {
                cancelAnimationFrame(frame);
                resizeObserver.disconnect();
                mutationObserver.disconnect();
            },
        };
    }

    $: displayImage = $questionImage || $currentQuestion?.questionImage || "";
    // export let game: HostGame;
</script>

<div class="page-background relative isolate flex h-dvh min-h-0 w-full flex-col overflow-hidden text-white">
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
        <div class="relative z-10 flex h-[20vh] min-h-28 max-h-56 shrink-0 items-center justify-center px-6 text-center">
            <div
                use:fitQuestionText
                class="rich-text-content w-full max-w-5xl text-6xl font-bold tracking-tight text-white drop-shadow"
            >
                {@html sanitizeRichText($currentQuestion.name)}
            </div>
        </div>

        <div class="relative z-10 flex min-h-0 flex-1 items-center shrink-0 justify-center p-4">
            {#if displayImage}
                <img
                    src={apiService.getImageUrl(displayImage)}
                    alt="Question visual"
                    class="max-h-full max-w-full rounded-2xl border border-white/20 object-contain shadow-2xl"
                />
            {/if}
        </div>

        <div class="relative z-10 flex h-[32vh] min-h-48 max-h-80 w-full shrink-0 flex-wrap">
            {#each COLORS as color, i}
                <QuizChoiceCard fixedHeight color={getCardColor($currentQuestion.choices[i], $gameState, color)}>
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