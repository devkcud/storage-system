<script lang="ts">
    import '$lib/styles/home.scss';

    import { onMount } from 'svelte';

    // @ts-ignore
    import FaTrash from 'svelte-icons/fa/FaTrash.svelte';
    // @ts-ignore
    import FaSearch from 'svelte-icons/fa/FaSearch.svelte';
    // @ts-ignore
    import FaPen from 'svelte-icons/fa/FaPen.svelte';
    // @ts-ignore
    import FaEye from 'svelte-icons/fa/FaEye.svelte';

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

{#if items.length > 0}
    <div id="products">
        {#each filteredItems as item}
            <div class="card">
                <div class="image">
                    <img src="https://placehold.co/150" alt="Product" />
                    <p>Product image</p>
                </div>
                <div class="content">
                    <header>
                        <h2>{item.name}</h2>
                    </header>

                    <div class="information">
                        <span class="quantity">Stock: <b>{item.quantity}</b></span>
                        <span class="price">Unit: <b>R${item.pricePerUnit.toFixed(2)}</b></span>
                    </div>

                    <p class="description">{item.description}</p>

                    {#if item.tags != null && item.tags.length > 0}
                        <ul class="tags">
                            {#each item.tags as tag}
                                <li>{tag}</li>
                            {/each}
                        </ul>
                    {/if}

                    <ul class="actions">
                        <li>
                            <a class="button" href="/item/{item.id}">
                                <FaEye />
                            </a>
                        </li>
                        <li>
                            <a class="button trash" href="/item/{item.id}/delete">
                                <FaTrash />
                            </a>
                        </li>
                        <li>
                            <a class="button edit" href="/item/{item.id}/edit">
                                <FaPen />
                            </a>
                        </li>
                    </ul>
                </div>
            </div>
        {/each}
    </div>
{:else if loading}
    <p class="center dimmed">Loading items...</p>
{:else}
    <p class="center dimmed">No items found</p>
{/if}
