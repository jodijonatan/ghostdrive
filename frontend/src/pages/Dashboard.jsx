import { useEffect, useState } from "react";
import api, { clearToken } from "../services/api";
import { saveFile, getFile } from "../db/indexeddb";

export default function Dashboard({ onLogout }) {
  const [files, setFiles] = useState([]);
  const [me, setMe] = useState(null);
  const [loading, setLoading] = useState(true);
  const [isUploading, setIsUploading] = useState(false);

  // ================= LOAD DATA =================
  const load = async () => {
    try {
      setLoading(true);
      const meRes = await api.get("/me");
      setMe(meRes.data);

      const fileRes = await api.get("/files");
      setFiles(Array.isArray(fileRes.data) ? fileRes.data : []);
    } catch (err) {
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    load();
  }, []);

  // ================= UPLOAD =================
  const upload = async (e) => {
    const file = e.target.files[0];
    if (!file) return;

    setIsUploading(true);
    try {
      const res = await api.post("/files", { filename: file.name });
      await saveFile(res.data.id, file);
      await load();
    } finally {
      setIsUploading(false);
    }
  };

  // ================= DOWNLOAD =================
  const download = async (id, filename) => {
    const file = await getFile(id);
    if (!file) return alert("File tidak ada di storage lokal!");

    const url = URL.createObjectURL(file);
    const a = document.createElement("a");
    a.href = url;
    a.download = filename;
    a.click();
  };

  // ================= DELETE =================
  const remove = async (id) => {
    if (!window.confirm("Hapus file ini?")) return;
    await api.delete(`/files/${id}`);
    await load();
  };

  // ================= RENAME =================
  const rename = async (id, oldName) => {
    const newName = prompt("Nama baru file:", oldName);
    if (!newName) return;

    await api.put(`/files/${id}`, { filename: newName });
    await load();
  };

  // ================= LOGOUT =================
  const logout = () => {
    clearToken();
    onLogout();
  };

  if (loading) {
    return (
      <div className="h-screen flex items-center justify-center">
        <p className="text-gray-500">Memuat dashboard...</p>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-50 p-8">
      {/* Header */}
      <div className="flex justify-between items-center mb-8">
        <div>
          <h1 className="text-2xl font-bold">Dashboard Drive</h1>
          <p className="text-sm text-gray-500">
            Login sebagai: {me?.email} ({me?.role})
          </p>
        </div>

        <div className="flex gap-3">
          <label className="bg-blue-600 text-white px-4 py-2 rounded cursor-pointer">
            {isUploading ? "Uploading..." : "Upload File"}
            <input
              type="file"
              className="hidden"
              onChange={upload}
              disabled={isUploading}
            />
          </label>

          <button
            onClick={logout}
            className="bg-red-500 text-white px-4 py-2 rounded"
          >
            Logout
          </button>
        </div>
      </div>

      {/* Table */}
      <div className="bg-white rounded shadow overflow-hidden">
        <table className="w-full text-left">
          <thead className="bg-gray-100 text-sm uppercase">
            <tr>
              <th className="p-4">Filename</th>
              {me?.role === "ADMIN" && <th className="p-4">Owner</th>}
              <th className="p-4 text-right">Aksi</th>
            </tr>
          </thead>
          <tbody>
            {files.length === 0 ? (
              <tr>
                <td colSpan="3" className="p-10 text-center text-gray-400">
                  Belum ada file
                </td>
              </tr>
            ) : (
              files.map((f) => (
                <tr key={f.id} className="border-t">
                  <td className="p-4 font-medium">{f.filename}</td>

                  {me?.role === "ADMIN" && (
                    <td className="p-4 text-sm text-gray-500">{f.owner_id}</td>
                  )}

                  <td className="p-4 text-right space-x-2">
                    <button
                      onClick={() => download(f.id, f.filename)}
                      className="text-blue-600"
                    >
                      Download
                    </button>
                    <button
                      onClick={() => rename(f.id, f.filename)}
                      className="text-yellow-600"
                    >
                      Rename
                    </button>
                    <button
                      onClick={() => remove(f.id)}
                      className="text-red-600"
                    >
                      Delete
                    </button>
                  </td>
                </tr>
              ))
            )}
          </tbody>
        </table>
      </div>
    </div>
  );
}
