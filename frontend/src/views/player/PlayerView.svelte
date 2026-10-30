<script lang="ts">
    import { GameState } from "../../service/net";
    import { PlayerGame, gameCancelled, gameState } from "../../service/player/player";
    import PlayerJoinView from "./PlayerJoinView.svelte";
    import PlayerLobbyView from "./PlayerLobbyView.svelte";
    import PlayerPlayView from "./PlayerPlayView.svelte";
    import PlayerRevealView from "./PlayerRevealView.svelte";
    import PlayerIntermissionView from "./PlayerIntermissionView.svelte";

    let game = new PlayerGame();
    let active = false;

    function onJoin() {
        active = true;
    }

    $: if ($gameCancelled) {
        active = false;
    }

    let views: Partial<Record<GameState, any>> = {
        [GameState.Lobby]: PlayerLobbyView,
        [GameState.Play]: PlayerPlayView,
        [GameState.Reveal]: PlayerRevealView,
        [GameState.Intermission]: PlayerIntermissionView
    };
</script>

{#if active}
    <svelte:component this={views[$gameState]} {game} />
{:else}
    <PlayerJoinView on:join={onJoin} {game} sessionClosed={$gameCancelled} />
{/if}
