export const saveFile = (id, file) => {
  return new Promise((resolve) => {
    const req = indexedDB.open("DriveMini", 1);
    req.onupgradeneeded = () => {
      req.result.createObjectStore("files");
    };
    req.onsuccess = () => {
      const db = req.result;
      const tx = db.transaction("files", "readwrite");
      tx.objectStore("files").put(file, id);
      resolve();
    };
  });
};

export const getFile = (id) => {
  return new Promise((resolve) => {
    const req = indexedDB.open("DriveMini", 1);
    req.onsuccess = () => {
      const db = req.result;
      const tx = db.transaction("files", "readonly");
      const getReq = tx.objectStore("files").get(id);
      getReq.onsuccess = () => resolve(getReq.result);
    };
  });
};
