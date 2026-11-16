<script lang="ts">
    import type { LeaderboardEntry } from "../service/net";
    import GlobalLiquidGlass from "./GlobalLiquidGlass.svelte";

    export let leaderboard: LeaderboardEntry[];
    export let finish = false;
    export let limit: number | undefined = undefined;
    export let position: number | undefined = undefined;

    const finishColors = ["bg-yellow-500", "bg-[silver]", "bg-yellow-700"];

    $: visibleLimit = limit !== undefined && Number.isFinite(limit) && limit > 0
        ? Math.floor(limit)
        : undefined;
    $: focusIndex = position !== undefined && Number.isFinite(position) && position >= 1
        ? Math.max(0, Math.min(leaderboard.length - 1, Math.floor(position) - 1))
        : undefined;
    $: visibleStart = visibleLimit === undefined || focusIndex === undefined
        ? 0
        : Math.max(0, Math.min(focusIndex - Math.floor(visibleLimit / 2), leaderboard.length - visibleLimit));
    $: visibleLeaderboard = visibleLimit === undefined
        ? leaderboard
        : leaderboard.slice(visibleStart, visibleStart + visibleLimit);
    $: scrollHeight = visibleLimit === undefined
        ? undefined
        : `min(${visibleLimit * 4.75}rem, 55dvh, calc(100dvh - 14rem))`;

    function autoScroll(node: HTMLDivElement, enabled: boolean) {
        let frame = 0;
        let previousTime: number | undefined;
        let direction = 1;

        function scroll(time: number) {
            if (enabled && !window.matchMedia("(prefers-reduced-motion: reduce)").matches
                && !node.matches(":hover")
                && !node.contains(document.activeElement)) {
                if (previousTime !== undefined) {
                    const maxScroll = node.scrollHeight - node.clientHeight;
                    if (maxScroll > 0) {
                        const nextPosition = node.scrollTop + direction * (time - previousTime) * 0.012;
                        if (nextPosition >= maxScroll) {
                            node.scrollTop = maxScroll;
                            direction = -1;
                        } else if (nextPosition <= 0) {
                            node.scrollTop = 0;
                            direction = 1;
                        } else {
                            node.scrollTop = nextPosition;
                        }
                    }
                }
            }

            previousTime = time;
            frame = requestAnimationFrame(scroll);
        }

        frame = requestAnimationFrame(scroll);

        return {
            update(nextEnabled: boolean) {
                enabled = nextEnabled;
                previousTime = undefined;
            },
            destroy() {
                cancelAnimationFrame(frame);
            },
        };
    }
</script>

<div class="w-full max-w-2xl rounded-3xl shadow-2xl">
    <GlobalLiquidGlass>
        <div class="flex flex-col gap-3 p-4 rounded-3xl">
            <h2 class="text-white text-center text-3xl">Leaderboard</h2>
            <!-- svelte-ignore a11y_no_noninteractive_tabindex -->
            <div
                class="flex flex-col gap-3 overflow-y-auto pr-2"
                style={scrollHeight ? `max-height: ${scrollHeight};` : undefined}
                use:autoScroll={visibleLimit !== undefined && focusIndex === undefined}
                role="region"
                aria-label="Player leaderboard"
                tabindex="0"
            >
                {#each visibleLeaderboard as entry, i}
                    {@const rank = visibleStart + i + 1}
                    {@const isFocused = focusIndex === rank - 1}
                    <div
                        class="flex min-h-16 items-center gap-4 rounded-xl p-2 text-xl text-white {isFocused ? 'bg-white/20 font-bold' : ''}"
                    >
                        {#if finish || focusIndex !== undefined}
                            <div
                                class="{rank <= 3 ? finishColors[rank - 1] : 'bg-transparent'} flex h-12 w-12 shrink-0 items-center justify-center rounded-full"
                            >
                                {rank}
                            </div>
                        {/if}
                        <p class="min-w-0 break-words">{entry.name} - {entry.points}</p>
                    </div>
                {/each}
            </div>
        </div>
    </GlobalLiquidGlass>
</div>
