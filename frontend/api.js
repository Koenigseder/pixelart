async function getPixels() {
  let result;

  await fetch("/api/pixels", {
    method: "GET",
  })
    .then((res) => res.json())
    .then((json) => (result = json));

  return result;
}
