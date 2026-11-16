<script lang="ts">
    export let color: string;
</script>

<div class="w-1/2 text-white text-center font-bold flex items-center justify-center text-3xl {color}">
    <slot />
</div>