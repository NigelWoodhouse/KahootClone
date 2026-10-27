<script lang="ts">
    import { createEventDispatcher } from "svelte";
    import { COLORS, type QuizQuestion } from "../../model/quiz";
    import Button from "../Button.svelte";
    import QuizChoiceCard from "../play/QuizChoiceCard.svelte";
    import Clock from "../Clock.svelte";
    import { apiService } from "../../service/api";

    const dispatch = createEventDispatcher();

    export let selectedQuestion: QuizQuestion;

    let fileInput: HTMLInputElement;
    let isUploading = false;
    let isDragOver = false;
    let showUrlInput = false;
    let manualUrl = "";

    function onDelete(){
        dispatch("delete");
    }

    function onTimeChange(e: Event){
        let target = e.target as HTMLInputElement;
        selectedQuestion.time = parseInt(target.value);
        dispatch("change");
    }

    function triggerFileInput() {
        if (fileInput) {
            fileInput.click();
        }
    }

    async function handleFile(file: File) {
        if (!file.type.startsWith("image/")) {
            alert("Please select a valid image file (PNG, JPG, GIF, WebP, etc.)");
            return;
        }

        isUploading = true;
        try {
            const result = await apiService.uploadImage(file);
            if (result && result.url) {
                selectedQuestion.image = result.url;
                dispatch("change");
            }
        } finally {
            isUploading = false;
        }
    }

    async function handleFileChange(e: Event) {
        const target = e.target as HTMLInputElement;
        if (target.files && target.files.length > 0) {
            await handleFile(target.files[0]);
            target.value = "";
        }
    }

    function handleDrop(e: DragEvent) {
        isDragOver = false;
        if (e.dataTransfer && e.dataTransfer.files && e.dataTransfer.files.length > 0) {
            handleFile(e.dataTransfer.files[0]);
        }
    }

    function removeImage() {
        selectedQuestion.image = "";
        manualUrl = "";
        dispatch("change");
    }

    function applyManualUrl() {
        if (manualUrl.trim()) {
            selectedQuestion.image = manualUrl.trim();
            showUrlInput = false;
            dispatch("change");
        }
    }
</script>

<div class="p-14 w-full flex-1">
	<div class="flex flex-col min-h-[95%] border rounded-lg shadow-sm bg-white overflow-hidden">
		<div
			class="flex font-bold text-center bg-gray-50 items-center border-b p-2 justify-between"
		>
			<div class="w-32"></div>
			<input
				class="p-4 text-3xl bg-gray-50 text-center font-bold outline-none border-b-2 border-transparent focus:border-purple-500 transition-colors"
                on:change
				bind:value={selectedQuestion.name}
			/>
			<div class="w-32 flex justify-end pr-2">
				<Button on:click={onDelete}>Delete</Button>
			</div>
		</div>

		<div class="flex-1 flex flex-col justify-center px-4 py-6">
			<div class="flex justify-between items-center gap-6">

				<Clock>
					<input
						value={selectedQuestion.time}
                        on:change={onTimeChange}
						type="text"
						class="w-[70%] text-base p-2 bg-purple-500 text-center text-white rounded font-bold"
					/>
				</Clock>

				<!-- Image Upload & Preview Section -->
				<div class="flex flex-col items-center flex-1 max-w-[560px]">
					<input
						type="file"
						accept="image/*"
						bind:this={fileInput}
						on:change={handleFileChange}
						class="hidden"
					/>

					{#if isUploading}
						<div class="w-full h-56 border-2 border-dashed border-purple-400 rounded-xl flex flex-col items-center justify-center bg-purple-50 text-purple-700 animate-pulse">
							<svg class="animate-spin h-8 w-8 mb-2 text-purple-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
								<circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
								<path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8H4z"></path>
							</svg>
							<span class="font-medium">Uploading image to MongoDB...</span>
						</div>
					{:else if selectedQuestion.image}
						<div class="relative group w-full flex flex-col items-center">
							<div class="relative max-h-56 max-w-full overflow-hidden rounded-xl border border-gray-200 shadow-md bg-gray-900 flex items-center justify-center">
								<img
									src={apiService.getImageUrl(selectedQuestion.image)}
									alt="Intermission preview"
									class="max-h-56 max-w-full object-contain"
								/>
								<div class="absolute inset-0 bg-black/40 opacity-0 group-hover:opacity-100 transition-opacity flex items-center justify-center gap-2">
									<button
										type="button"
										on:click={triggerFileInput}
										class="px-3 py-1.5 bg-white text-gray-800 text-xs font-semibold rounded-md shadow hover:bg-gray-100 transition"
									>
										Replace
									</button>
									<button
										type="button"
										on:click={removeImage}
										class="px-3 py-1.5 bg-red-600 text-white text-xs font-semibold rounded-md shadow hover:bg-red-700 transition"
									>
										Remove
									</button>
								</div>
							</div>
							<div class="mt-1 flex items-center justify-between w-full text-xs text-gray-500 px-1">
								<span class="flex items-center gap-1 text-emerald-600 font-medium">
									<svg class="w-3.5 h-3.5 inline" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" /></svg>
									Intermission image saved
								</span>
								<button type="button" on:click={removeImage} class="text-red-500 hover:underline">Clear</button>
							</div>
						</div>
					{:else}
						<button
							type="button"
							class="w-full h-52 border-2 border-dashed rounded-xl flex flex-col items-center justify-center cursor-pointer transition-all {isDragOver ? 'border-purple-500 bg-purple-50 scale-102' : 'border-gray-300 hover:border-purple-400 hover:bg-gray-50'}"
							on:click={triggerFileInput}
							on:dragover|preventDefault={() => isDragOver = true}
							on:dragleave={() => isDragOver = false}
							on:drop|preventDefault={handleDrop}
						>
							<div class="p-3 bg-purple-100 text-purple-600 rounded-full mb-2">
								<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
									<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
								</svg>
							</div>
							<span class="font-semibold text-gray-700 text-sm">Upload Intermission Image</span>
							<span class="text-xs text-gray-500 mt-1">Click to browse or drag & drop (saved to MongoDB)</span>
							<span class="text-[11px] text-purple-600 font-medium mt-1">Displayed during IntermissionState</span>
						</button>
						
						{#if !showUrlInput}
							<button
								type="button"
								on:click={() => showUrlInput = true}
								class="text-xs text-purple-600 hover:underline mt-1.5"
							>
								or enter image URL
							</button>
						{:else}
							<div class="flex gap-1 mt-2 w-full">
								<input
									type="text"
									placeholder="https://example.com/image.png"
									bind:value={manualUrl}
									class="flex-1 text-xs border rounded px-2 py-1"
								/>
								<button
									type="button"
									on:click={applyManualUrl}
									class="px-2 py-1 bg-purple-600 text-white rounded text-xs"
								>
									Set
								</button>
								<button
									type="button"
									on:click={() => showUrlInput = false}
									class="px-2 py-1 bg-gray-200 text-gray-600 rounded text-xs"
								>
									Cancel
								</button>
							</div>
						{/if}
					{/if}
				</div>

				<div class="w-24"></div>

			</div>
		</div>

		<div class="flex flex-wrap w-full h-96">
			{#each COLORS as color, i}
				<QuizChoiceCard {color}>
					<div class="px-14 w-full flex gap-2">
						<input
							class="rounded px-2 py-1 w-full text-black"
							placeholder="Choice"
							bind:value={selectedQuestion.choices[i].name}
                            on:input={() => dispatch("change")}
						/>
						<input
							type="checkbox"
							class="w-16 h-18 cursor-pointer"
							bind:checked={selectedQuestion.choices[i].correct}
                            on:change={() => dispatch("change")}
						/>
					</div>
				</QuizChoiceCard>
			{/each}
		</div>
	</div>
</div>
