<script lang="ts">
    import '$lib/styles/home.scss';

    import { onMount } from 'svelte';
    import { writable } from 'svelte/store';

    // @ts-ignore
    import FaTrash from 'svelte-icons/fa/FaTrash.svelte';
    // @ts-ignore
    import FaPen from 'svelte-icons/fa/FaPen.svelte';

    let items: {
        id: string;
        name: string;
        description: string;
        pricePerUnit: number;
        quantity: number;
        tags: string[];
    }[];

    let loading = true;

    onMount(async () => {
        const response = await fetch('http://localhost:8080/', { method: 'GET' });
        items = await response.json();

        loading = false;
    });
</script>

<h1 id="title">Products</h1>

{#if items}
    <div id="products">
        {#each items as item}
            <div class="card">
                <header>
                    <h2>{item.name}</h2>

                    <div>
                        <a class="edit" href="/item/{item.id}/edit">
                            <FaPen />
                        </a>
                        <a class="trash" href="/item/{item.id}/delete">
                            <FaTrash />
                        </a>
                    </div>
                </header>

                <div class="information">
                    <span class="quantity">Stock: <b>{item.quantity}</b></span>
                    <span class="price">Unit: <b>R${item.pricePerUnit.toFixed(2)}</b></span>
                </div>

                <p>{item.description}</p>

                {#if item.tags != null && item.tags.length > 0}
                    <ul>
                        {#each item.tags as tag}
                            <li>{tag}</li>
                        {/each}
                    </ul>
                {/if}

                <a href="/item/{item.id}">Manage product</a>
            </div>
        {/each}
    </div>
{:else if loading}
    <p class="center dimmed">Loading items...</p>
{:else}
    <p class="center dimmed">Ain't nobody here but us chickens (no items found)</p>
{/if}
