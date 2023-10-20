<script lang="ts">
    import '$lib/styles/item.scss';

    import { goto } from '$app/navigation';

    // @ts-ignore
    import FaArrowLeft from 'svelte-icons/fa/FaArrowLeft.svelte';
</script>

<button id="go-back" on:click={() => goto('/')}>
    <FaArrowLeft />
</button>

<main id="item-holder">
    <slot />
</main>
