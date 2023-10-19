<script lang="ts">
    import '$lib/styles/home.scss';

    import { onMount } from 'svelte';

    // @ts-ignore
    import FaTrash from 'svelte-icons/fa/FaTrash.svelte';
    // @ts-ignore
    import FaSearch from 'svelte-icons/fa/FaSearch.svelte';
    // @ts-ignore
    import FaPen from 'svelte-icons/fa/FaPen.svelte';

    interface Item {
        id: string;
        name: string;
        description: string;
        pricePerUnit: number;
        quantity: number;
        tags: string[] | null;
    }

    let items: Item[] = [];
    let filteredItems: Item[] = [];
    let loading = true;
    let searchQuery = '';

    onMount(async () => {
        try {
            const response = await fetch('http://localhost:8080/', { method: 'GET' });
            items = await response.json();
            filteredItems = items;
        } catch (error) {
            console.error('Failed to fetch items:', error);
        } finally {
            loading = false;
        }

        filteredItems = items;
    });

    function filterItems() {
        if (searchQuery === '') {
            filteredItems = items;
        } else {
            filteredItems = items.filter((item) => {
                const nameMatch = item.name.toLowerCase().includes(searchQuery.toLowerCase());
                const descriptionMatch = item.description
                    .toLowerCase()
                    .includes(searchQuery.toLowerCase());
                const tagMatch =
                    item.tags &&
                    item.tags.some((tag) => tag.toLowerCase().includes(searchQuery.toLowerCase()));
                return nameMatch || descriptionMatch || tagMatch;
            });
        }
    }
</script>

<h1 id="title">Products</h1>

<div id="search">
    <FaSearch />

    <input
        type="text"
        id="search"
        bind:value={searchQuery}
        placeholder="Search name, description or tag"
        on:input={filterItems}
    />
</div>

{#if filteredItems.length > 0}
    <div id="products">
        {#each filteredItems as item}
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
    <p class="center dimmed">No items found</p>
{/if}
