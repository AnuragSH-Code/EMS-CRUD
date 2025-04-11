// used production by default for now .
const BASE_URL =
  import.meta.env.MODE === "development"
    ? import.meta.env.VITE_BACKEND_DEVELOPMENT_URL
    : import.meta.env.VITE_BACKEND_PRODUCTION_URL;

export async function apiFetch(endpoint, options = {}) {
  const res = await fetch(`${BASE_URL}${endpoint}`, {
    headers: { "Content-Type": "application/json", ...options.headers },
    ...options,
  });

  if (!res.ok) throw new Error(`API Error: ${res.status}`);
  return res.json();
}
