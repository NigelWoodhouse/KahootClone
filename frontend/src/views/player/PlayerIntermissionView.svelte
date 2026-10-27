<script lang="ts">
    import { intermissionQuestionName, intermissionImage, leaderboard, points, type PlayerGame } from "../../service/player/player";
    import { apiService } from "../../service/api";
    import Leaderboard from "../../lib/Leaderboard.svelte";

    export let game: PlayerGame;
</script>

<div class="min-h-screen w-full bg-gradient-to-br from-indigo-900 via-purple-900 to-indigo-950 text-white flex flex-col items-center justify-center p-6">
    <div class="w-full max-w-md flex flex-col items-center text-center">
        <span class="bg-yellow-400 text-purple-950 px-4 py-1 rounded-full text-xs font-black tracking-widest uppercase mb-4 shadow">
            Intermission
        </span>

        {#if $intermissionQuestionName}
            <h2 class="text-xl font-bold mb-3 text-white/90">
                {$intermissionQuestionName}
            </h2>
        {/if}

        {#if $intermissionImage}
            <div class="w-full bg-white/10 backdrop-blur-md p-3 rounded-2xl border border-white/20 shadow-xl mb-6 overflow-hidden flex items-center justify-center">
                <img
                    src={apiService.getImageUrl($intermissionImage)}
                    alt="Question visual"
                    class="max-h-60 max-w-full object-contain rounded-xl shadow"
                />
            </div>
        {/if}

        {#if $leaderboard && $leaderboard.length > 0}
            <div class="w-full bg-white/10 backdrop-blur-md p-5 rounded-2xl border border-white/15 shadow-lg mb-4">
                <h3 class="text-sm font-bold text-yellow-300 uppercase tracking-wider mb-3">Leaderboard</h3>
                <Leaderboard leaderboard={$leaderboard} />
            </div>
        {:else}
            <div class="bg-white/10 backdrop-blur-md p-6 rounded-2xl border border-white/15 shadow-lg mb-4 w-full">
                <p class="text-3xl font-extrabold text-yellow-300">{$points} pts</p>
                <p class="text-sm text-white/70 mt-1">Get ready for the next question!</p>
            </div>
        {/if}

        <p class="text-xs text-white/50 tracking-wide mt-2 animate-pulse">
            Waiting for next question...
        </p>
    </div>
</div>
