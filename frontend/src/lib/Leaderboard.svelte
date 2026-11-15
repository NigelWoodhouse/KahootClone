<script lang="ts">
    import type { LeaderboardEntry } from "../service/net";
    import GlobalLiquidGlass from "./GlobalLiquidGlass.svelte";

    export let leaderboard: LeaderboardEntry[];
    export let finish = false;

    const finishColors = ["bg-yellow-500", "bg-[silver]", "bg-yellow-700", "bg-transparent"];
</script>

    <div class="flex rounded-xl p-4 flex-col gap-4 w-full max-w-2xl">
    <GlobalLiquidGlass>
        <h2 class="text-white text-center text-3xl">Leaderboard</h2>
        {#each leaderboard as entry, i}
            <div
                class="text-white p-2 text-2xl rounded-xl flex items-center gap-6"
            >
                {#if finish}
                    <div
                        class="{finishColors[
                            i < 3 ? i : 4
                        ]} w-14 h-14 rounded-full flex items-center justify-center ml-8"
                    >
                        {i + 1}
                    </div>
                {/if}
                <p class="min-w-0 break-words">{entry.name} - {entry.points}</p>
            </div>
        {/each}
        </GlobalLiquidGlass>
    </div>

