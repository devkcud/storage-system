<script lang="ts">
    import { onMount } from 'svelte';
    import { page } from '$app/stores';

    // @ts-ignore
    import FaTrash from 'svelte-icons/fa/FaTrash.svelte';
    // @ts-ignore
    import FaPen from 'svelte-icons/fa/FaPen.svelte';

    interface Item {
        name: string;
        description: string;
        pricePerUnit: number;
        quantity: number;
        tags: string[] | null;
    }

    let item: Item | undefined;
    let loading = true;

    onMount(async () => {
        try {
            const response = await fetch(`http://localhost:8080/${$page.params.id}`, {
                method: 'GET',
            });

            if (response.ok) {
                item = await response.json();
            }
        } catch (error) {
            console.error('Failed to fetch item:', error);
        } finally {
            loading = false;
        }
    });
</script>

{#if item}
    <div class="content">
        <header>
            <h1>{item.name}</h1>

            <div class="information">
                <span class="quantity">Stock: <b>{item.quantity}</b></span>
                <span class="price">Unit: <b>R${item.pricePerUnit.toFixed(2)}</b></span>
                <span class="price">
                    Total: <b>R${(item.pricePerUnit * item.quantity).toFixed(2)}</b>
                </span>
            </div>
        </header>

        <div class="description">
            <h2>Description</h2>
            <p class="description">{item.description}</p>
        </div>

        <div class="tags">
            {#if item.tags != null && item.tags.length > 0}
                {#each item.tags as tag}
                    <p>{tag}</p>
                {/each}
            {:else}
                <p class="no-tags">No tags</p>
            {/if}
        </div>

        <div class="actions">
            <a class="button trash" href="/item/{$page.params.id}/delete">
                <FaTrash />
                Delete
            </a>
            <a class="button" href="/item/{$page.params.id}/edit">
                <FaPen />
                Edit
            </a>
        </div>
    </div>
{:else if loading}
    <p>Loading item...</p>
{:else}
    <p>Couldn't fetch item with ID {$page.params.id}</p>
{/if}
