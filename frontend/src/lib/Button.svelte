<button on:click class="button-background bg-blue-500 hover:bg-blue-600 text-white rounded-md px-4 py-2">
    <div>
        <slot />
    </div>
</button>

<style>
    .button-background {
        background: linear-gradient(to bottom, #b6e8ae, #5cd44a);
    }
</style>