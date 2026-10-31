<script lang="ts">
    import type { PlayerGame } from "../../service/player/player";

    export let game: PlayerGame;

    const circles = [
        { left: "25%", size: 100, delay: 0, duration: 25, },
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

<div
    class="page-background relative isolate flex min-h-screen w-full items-center justify-center overflow-hidden px-6 py-12 text-center"
>
    <ul class="bubbles" aria-hidden="true">
        {#each circles as circle}
            <li
                style="left: {circle.left}; width: {circle.size}px; height: {circle.size}px; animation-delay: {circle.delay}s; animation-duration: {circle.duration}s"
            ></li>
        {/each}
    </ul>

    <main class="relative z-10 space-y-3 text-white">
        <h1 class="text-4xl font-bold drop-shadow-md sm:text-5xl">
            Welcome to the game!
        </h1>
        <p class="text-lg sm:text-xl">Do you see your name on the screen?</p>
    </main>
</div>

<style>
    .bubbles {
        position: absolute;
        inset: 0;
        overflow: hidden;
        margin: 0;
        padding: 0;
        list-style: none;
        pointer-events: none;
    }

    .bubbles li {
        position: absolute;
        bottom: -150px;
        display: block;
        background: rgb(255 255 255 / 20%);
        animation: float-up 25s linear infinite;
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
        .bubbles li {
            animation: none;
        }
    }
</style>