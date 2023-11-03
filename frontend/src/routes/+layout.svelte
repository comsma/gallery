<script>
    import "../app.css"
</script>

<div class="flex pt-10 mx-5">
    <aside class="h-screen sticky top-0">
        <div class="max-w-[800px]">
            <img src="/COMSMA.webp">
        </div>
        <div class="text-sm text-slate-600">
            <a href="https://comsma.com">about me</a>
        </div>

    </aside>
    <main class="ml-5">
        <slot />
    </main>

</div>

<style>
    a{
        font-family: "Pier Sans";
    }
</style>
