async function dataUrlToFile(dataUrl, fileName) {
    const res = await fetch(dataUrl);
    const blob  = await res.blob();
    return new File([blob], fileName, { type: dataUrl.match(/^data:(.+);base64/)?.[1] });
}