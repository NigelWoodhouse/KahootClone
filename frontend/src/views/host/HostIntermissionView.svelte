<script lang="ts">
    import Button from "../../lib/Button.svelte";
    import Leaderboard from "../../lib/Leaderboard.svelte";
    import { HostGame, intermissionImage, leaderboard, intermissionQuestionName, currentQuestion, tick } from "../../service/host/host";
    import { apiService } from "../../service/api";

    export let game: HostGame;

    function skip() {
        game.start();
    }

    $: displayImage = $intermissionImage || $currentQuestion?.image || "";
    $: questionName = $intermissionQuestionName || $currentQuestion?.name || "";
</script>

<div class="bg-gradient-to-br from-purple-900 via-indigo-900 to-purple-800 min-h-screen w-full flex flex-col text-white">
    <!-- Top Header Bar -->
    <div class="flex justify-between items-center px-10 py-6 border-b border-white/10 bg-black/20 backdrop-blur-sm">
        <div class="flex items-center gap-4">
            <span class="bg-yellow-400 text-purple-950 px-3.5 py-1 rounded-full text-sm font-extrabold tracking-wider uppercase shadow-md">
                Intermission
            </span>
            {#if questionName}
                <h1 class="text-2xl font-bold tracking-tight text-white/90 drop-shadow">
                    {questionName}
                </h1>
            {/if}
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
    <div class="flex-1 flex items-center justify-center p-8">
        {#if displayImage}
            <div class="w-full max-w-6xl grid grid-cols-1 lg:grid-cols-2 gap-10 items-center">
                <!-- Intermission Image Card -->
                <div class="flex flex-col items-center">
                    <div class="bg-white/10 backdrop-blur-md p-4 rounded-3xl border border-white/20 shadow-2xl overflow-hidden max-h-[500px] flex items-center justify-center transition-all hover:scale-[1.01]">
                        <img
                            src={apiService.getImageUrl(displayImage)}
                            alt="Question intermission visual"
                            class="max-h-[460px] w-auto max-w-full object-contain rounded-2xl shadow-inner"
                        />
                    </div>
                    {#if questionName}
                        <p class="mt-3 text-sm font-medium text-white/70 italic text-center">
                            "{questionName}"
                        </p>
                    {/if}
                </div>

                <!-- Leaderboard Section -->
                <div class="flex flex-col items-center">
                    <div class="w-full max-w-md bg-white/10 backdrop-blur-md p-6 rounded-3xl border border-white/20 shadow-2xl">
                        <h2 class="text-xl font-bold text-center text-yellow-300 mb-4 tracking-wide uppercase">
                            Scoreboard
                        </h2>
                        <Leaderboard leaderboard={$leaderboard} />
                    </div>
                </div>
            </div>
        {:else}
            <!-- If no image, display centered leaderboard -->
            <div class="flex flex-col items-center">
                <div class="w-full max-w-md bg-white/10 backdrop-blur-md p-8 rounded-3xl border border-white/20 shadow-2xl">
                    <h2 class="text-2xl font-extrabold text-center text-yellow-300 mb-6 tracking-wide uppercase">
                        Scoreboard
                    </h2>
                    <Leaderboard leaderboard={$leaderboard} />
                </div>
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
