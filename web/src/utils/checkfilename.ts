function getFileExtension(filename: string) {
  const lastDotIndex = filename.lastIndexOf(".");
  if (lastDotIndex === -1) {
    return "";
  }
  return filename.slice(lastDotIndex + 1).toLowerCase();
}

export function isVideo(filename: string) {
  const ext = getFileExtension(filename);
  const videoExtensions = ["mp4", "mov", "avi", "wmv", "flv", "mkv"];
  return videoExtensions.includes(ext);
}

export function isImage(filename: string) {
  const ext = getFileExtension(filename);
  const imageExtensions = ["jpg", "jpeg", "png", "gif", "bmp", "webp"];
  return imageExtensions.includes(ext);
}
