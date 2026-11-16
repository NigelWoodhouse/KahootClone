<script lang="ts">
    import GlobalLiquidGlass from "./GlobalLiquidGlass.svelte";
</script>

<div class="rounded-xl">
    <GlobalLiquidGlass>
        <button
            on:click
            class="inline-flex items-center justify-center rounded-xl border border-white/20 px-5 py-2 text-white transition-all duration-200 hover:-translate-y-0.5 hover:border-white/40 hover:bg-white/20 hover:shadow-lg active:translate-y-0 active:scale-95 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-white"
        >
            <slot />
        </button>
    </GlobalLiquidGlass>
</div>

<style>

</style>