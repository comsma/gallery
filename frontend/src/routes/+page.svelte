<script lang="ts">
const images = [
    {
        src: "https://comsma-gallery.nyc3.cdn.digitaloceanspaces.com/DJI_0899.jpg"
    },
    {
        src: "https://comsma-gallery.nyc3.cdn.digitaloceanspaces.com/DJI_0093-HDR.jpg"
    },
    {
        src: "https://comsma-gallery.nyc3.cdn.digitaloceanspaces.com/DJI_0034.jpg"
    },
    {
        src: "https://comsma-gallery.nyc3.cdn.digitaloceanspaces.com/DJI_0079.jpg"
    },
    {
        src: "https://comsma-gallery.nyc3.cdn.digitaloceanspaces.com/DJI_0315.jpg"
    },
    {
        src: "https://comsma-gallery.nyc3.cdn.digitaloceanspaces.com/0C6A8291.jpg"
    },
    {
        src: "https://comsma-gallery.nyc3.cdn.digitaloceanspaces.com/0C6A9699.jpg"
    },
    {
        src: "https://comsma-gallery.nyc3.cdn.digitaloceanspaces.com/Untitled%20(3)%20(2).jpg"
    },
    {
        src: "https://comsma-gallery.nyc3.cdn.digitaloceanspaces.com/DJI_0254-HDR.jpg"
    },
    {
        src: "https://comsma-gallery.nyc3.cdn.digitaloceanspaces.com/0C6A7926.jpg"
    },
    {
        src: "https://comsma-gallery.nyc3.cdn.digitaloceanspaces.com/_C6A6000.jpg"
    },
    {
        src: "https://comsma-gallery.nyc3.cdn.digitaloceanspaces.com/DJI_0054.jpg"
    },
    {
        src: "https://comsma-gallery.nyc3.cdn.digitaloceanspaces.com/_C6A6084.jpg"
    },{
        src: "https://comsma-gallery.nyc3.cdn.digitaloceanspaces.com/DJI_0004.jpg"
    },{
        src: "https://comsma-gallery.nyc3.cdn.digitaloceanspaces.com/DJI_0070.jpg"
    },{
        src: "https://comsma-gallery.nyc3.cdn.digitaloceanspaces.com/IMG_0503.jpg"
    },{
        src: "https://comsma-gallery.nyc3.cdn.digitaloceanspaces.com/f9820416.jpg"
    },

]
</script>

<div class="grid grid-cols-1 gap-4 lg:grid-cols-3">
    {#each images as image}
        <div class="object-cover">
            <img alt="picture" src={image.src} />
        </div>

    {/each}
</div>

<style>
    :global(body) {
        background-color: #f3f1f2;
    }
</style>
