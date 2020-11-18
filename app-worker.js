const cacheName = "app-" + "997eccc74d75987c8dacd060c5864e0f4bb751ac";

self.addEventListener("install", event => {
  console.log("installing app worker 997eccc74d75987c8dacd060c5864e0f4bb751ac");
  self.skipWaiting();

  event.waitUntil(
    caches.open(cacheName).then(cache => {
      return cache.addAll([
        "/",
        "/app.css",
        "/app.js",
        "/manifest.json",
        "/wasm_exec.js",
        "/web/app.wasm",
        "/web/static/notes.css",
        "http://cdn.materialdesignicons.com/5.4.55/css/materialdesignicons.min.css",
        "https://i.imgur.com/vNxAhoY.png",
        "https://www.w3schools.com/w3css/4/w3pro.css",
        
      ]);
    })
  );
});

self.addEventListener("activate", event => {
  event.waitUntil(
    caches.keys().then(keyList => {
      return Promise.all(
        keyList.map(key => {
          if (key !== cacheName) {
            return caches.delete(key);
          }
        })
      );
    })
  );
  console.log("app worker 997eccc74d75987c8dacd060c5864e0f4bb751ac is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
