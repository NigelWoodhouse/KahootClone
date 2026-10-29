<script lang="ts">
    import Button from "../../lib/Button.svelte";
    import PlayerNameCard from "../../lib/lobby/PlayerNameCard.svelte";
    import { players, type HostGame, gameCode } from "../../service/host/host";

    export let game: HostGame;

    function start() {
        game.start();
    }

    const bubblePatterns = [
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

    function getBubblePattern(index: number) {
        const pattern = bubblePatterns[index % bubblePatterns.length];
        if (index < bubblePatterns.length) {
            return pattern;
        }

        return {
            ...pattern,
            left: `${5 + ((index * 37) % 85)}%`,
            size: 100,
            duration: 18 + (index % 5) * 2,
        };
    }

    $: bubbles = Array.from(
        { length: Math.max(bubblePatterns.length, $players.length) },
        (_, index) => {
            const pattern = getBubblePattern(index);
            const player = $players[index];

            return {
                ...pattern,
                player,
                animationDelay: player ? -pattern.duration * 0.45 : pattern.delay,
            };
        },
    );
</script>

<div class="relative isolate min-h-screen w-full overflow-hidden p-8 lobby-background">
    <ul class="bubbles" aria-hidden="true">
        {#each bubbles as bubble, i (bubble.player?.id ?? `decoration-${i}`)}
            <li
                style="left: {bubble.left}; width: {bubble.size}px; height: {bubble.size}px; animation-delay: {bubble.animationDelay}s; animation-duration: {bubble.duration}s"
            >
                {#if bubble.player}
                    <span>{bubble.player.name}</span>
                {/if}
            </li>
        {/each}
    </ul>

    <div class="relative z-10">
        <div class="flex justify-end">
            <Button on:click={start}>Start game</Button>
        </div>
        <div class="text-center text-white">
            <h2 class="text-4xl">Join with game code</h2>
            <h2 class="mt-4 text-6xl font-bold">{$gameCode}</h2>
        </div>
        <!-- <h2 class="mt-10 text-4xl text-center text-white font-bold text-white">
            Players ({$players.length})
        </h2> -->
        <!-- <div class="mt-4 flex flex-wrap gap-2">
            {#each $players as player (player.id)}
                <PlayerNameCard {player} />
            {:else}
                <p class="text-white">No players have joined yet</p>
            {/each}
        </div> -->
        <div class="mt-10 text-center text-white">
            {#if $players.length == 0}
                <p class="text-white">No players have joined yet</p>
            {:else if $players.length == 1}
                <p class="mt-10 text-center text-white text-white">
                    {$players.length} player is in the lobby
                </p>
            {:else}
                <p class="mt-10 text-center text-white text-white">
                    {$players.length} players are in the lobby
                </p>
            {/if}
        </div>
    </div>
</div>

<style>
    .lobby-background {
        background: linear-gradient(to left, #8f94fb, #4e54c8);
    }

    .bubbles {
        position: absolute;
        inset: 0;
        overflow: hidden;
        margin: 0;
        padding: 0;
        list-style: none;
        pointer-events: none;
        z-index: 0;
    }

    .bubbles li {
        position: absolute;
        bottom: -150px;
        display: flex;
        align-items: center;
        justify-content: center;
        background: rgb(255 255 255 / 20%);
        animation: float-up 25s linear infinite;
    }

    .bubbles li span {
        max-width: 80%;
        color: white;
        font-size: clamp(0.7rem, 1.5vw, 1rem);
        font-weight: 700;
        line-height: 1.2;
        overflow-wrap: anywhere;
        text-align: center;
        text-shadow: 0 1px 3px rgb(0 0 0 / 40%);
    }

    @keyframes float-up {
        0% {
            transform: translateY(0) rotate(0deg);
            opacity: 1;
            border-radius: 0;
        }

        100% {
            transform: translateY(calc(-100vh - 150px)) rotate(720deg);
            opacity: 0;
            border-radius: 50%;
        }
    }

    @media (prefers-reduced-motion: reduce) {
        .bubbles li,
        .bubbles li span {
            animation-play-state: paused;
        }
    }
</style>
