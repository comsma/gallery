<script>
    import "../app.css"
</script>

<div class="flex py-10 mx-5 flex-col lg:flex-row">
    <div class="block lg:hidden">
        <div class="mx-auto object-center">
            <img height="auto" width="300px" class="m-auto" alt="comsma logo" src="/COMSMA.webp">
        </div>
    </div>
    <aside class="hidden h-screen sticky top-0 lg:block">
        <div class="">
            <img height="auto" width="1000px" alt="comsma logo" src="/COMSMA.webp">
        </div>
        <div class="text-sm text-slate-600">
            <a href="https://comsma.com">about me</a>
        </div>

    </aside>
    <main class="lg:ml-5">
        <slot />
    </main>

</div>

<style>
    a{
        font-family: "Pier Sans";
    }
</style>
