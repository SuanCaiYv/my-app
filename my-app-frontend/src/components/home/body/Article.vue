<template>
    <div class="article">
        <div class="articleHead">
            {{title}}
        </div>
        <div class="articleBody">
            {{body}}
        </div>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<String>("Article")

let props = defineProps({
    title: String,
    body: String
})
</script>

<style scoped>
.article {
}

.articleHead {
    height: 40px;
    width: 320px;
    margin: 0 auto 0 50px;
    border: 1px solid black;
    box-sizing: border-box;
    padding: 0;
}

.articleBody {
    height: 200px;
    width: auto;
    margin: 0 50px 20px;
    border: 1px solid black;
    box-sizing: border-box;
    padding: 0;
}
</style>