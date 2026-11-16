<script lang="ts">
    import { intermissionQuestionName, intermissionImage, leaderboard, playerName, totalPoints, type PlayerGame } from "../../service/player/player";
    import { apiService } from "../../service/api";
    import Leaderboard from "../../lib/Leaderboard.svelte";

    export let game: PlayerGame;

    $: matchedPlayerIndex = $leaderboard.findIndex(
        (entry) => entry.name === $playerName && entry.points === $totalPoints,
    );
    $: playerIndex = matchedPlayerIndex >= 0
        ? matchedPlayerIndex
        : $leaderboard.findIndex((entry) => entry.name === $playerName);
    $: currentPosition = playerIndex >= 0 ? playerIndex + 1 : undefined;

    const bubbles = [
        { left: "25%", size: 100, delay: 0, duration: 25 },
        { left: "10%", size: 40, delay: 2, duration: 12 },
        { left: "70%", size: 50, delay: 4, duration: 25 },
        { left: "40%", size: 80, delay: 0, duration: 18 },
        { left: "65%", size: 60, delay: 3, duration: 25 },
        { left: "75%", size: 130, delay: 7, duration: 25 },
        { left: "35%", size: 150, delay: 15, duration: 25 },
        { left: "50%", size: 45, delay: 3, duration: 45 },
        { left: "20%", size: 35, delay: 2, duration: 35 },
        { left: "85%", size: 170, delay: 0, duration: 11 },
    ];
</script>

<div class="page-background relative isolate min-h-screen w-full overflow-hidden text-white flex flex-col items-center justify-center p-6">
    <ul class="bubbles" aria-hidden="true">
        {#each bubbles as bubble}
            <li
                style="left: {bubble.left}; width: {bubble.size}px; height: {bubble.size}px; animation-delay: {bubble.delay}s; animation-duration: {bubble.duration}s"
            ></li>
        {/each}
    </ul>
    <div class="relative z-10 w-full max-w-md flex flex-col items-center text-center">
        <span class="bg-yellow-400 text-purple-950 px-4 py-1 rounded-full text-xs font-black tracking-widest uppercase mb-4 shadow">
            Intermission
        </span>

        <!-- {#if $intermissionQuestionName}
            <h2 class="text-xl font-bold mb-3 text-white/90">
                {$intermissionQuestionName}
            </h2>
        {/if} -->

        <!-- {#if $intermissionImage}
            <div class="w-full bg-white/10 backdrop-blur-md p-3 rounded-2xl border border-white/20 shadow-xl mb-6 overflow-hidden flex items-center justify-center">
                <img
                    src={apiService.getImageUrl($intermissionImage)}
                    alt="Question visual"
                    class="max-h-60 max-w-full object-contain rounded-xl shadow"
                />
            </div>
        {/if} -->

        <!-- <div class="w-full rounded-2xl bg-purple-600 px-5 py-3 shadow-lg mb-4">
            <p class="text-sm text-white/75">Your total score</p>
            <p class="text-3xl font-extrabold text-yellow-300">{$totalPoints} pts</p>
        </div> -->

        {#if $leaderboard && $leaderboard.length > 0}
            <!-- <div class="w-full backdrop-blur-md p-5 rounded-2xl border shadow-lg mb-4"> -->
                <!-- <h3 class="text-sm font-bold text-yellow-300 uppercase tracking-wider mb-3">Leaderboard</h3> -->
                <Leaderboard
                    leaderboard={$leaderboard}
                    limit={3}
                    position={currentPosition}
                />
            <!-- </div> -->
        {:else}
            <div class="bg-white/10 backdrop-blur-md p-6 rounded-2xl border border-white/15 shadow-lg mb-4 w-full">
                <p class="text-sm text-white/70 mt-1">Get ready for the next question!</p>
            </div>
        {/if}

        <p class="text-xs text-white/50 tracking-wide mt-2 animate-pulse">
            Waiting for next question...
        </p>
    </div>
</div>
