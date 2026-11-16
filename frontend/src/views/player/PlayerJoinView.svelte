<script lang="ts">
    import { onMount } from "svelte";
    import { createEventDispatcher } from "svelte";
    import Button from "../../lib/Button.svelte";
    import type { PlayerGame } from "../../service/player/player";

    const dispatch = createEventDispatcher();

    let code: string = "";
    let name: string = "";
    export let game: PlayerGame;
    export let sessionClosed = false;

    function join(){
        dispatch("join");
        game.join(code, name);
    }

    onMount(() => {
        const params = new URLSearchParams(window.location.search);
        const debugCode = params.get("code");
        const debugName = params.get("name");
        if (params.get("debug") !== "players" || !debugCode || !debugName) {
            return;
        }

        code = debugCode;
        name = debugName;
        join();
    });
</script>

<div class="page-background min-h-screen w-full flex items-center justify-center">
    <div>
        <h2 class="text-white font-bold text-5xl text-center">Quiz</h2>
        {#if sessionClosed}
            <p class="mt-4 text-center text-white" role="status">
                The host has closed this session. Join another game when you are ready.
            </p>
        {/if}
        <div class="flex flex-col gap-2 mt-10 items-center rounded-xl">
            <input bind:value={code} type="text" placeholder="Game Code" class="p-2 rounded" />
            <input bind:value={name} type="text" placeholder="Name" class="p-2 rounded" />
            <Button on:click={join}>Join Game</Button>
        </div>
    </div>
</div>
