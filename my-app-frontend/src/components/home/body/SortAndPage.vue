<template>
    <div class="sortAndPage">
        aaa
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<String>("SortAndPage")
</script>

<style scoped>
.sortAndPage {
    width: 100%;
    height: 100%;
    grid-area: sortAndPage;
    border: 1px solid lightgrey;
    box-sizing: border-box;
}
</style>