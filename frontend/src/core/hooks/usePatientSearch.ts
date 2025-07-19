import { useEffect, useState } from "react";
import axios from "axios";

export function usePatientSearch(query: string) {
  const [results, setResults] = useState<any[]>([]);
  const [loading, setLoading] = useState(false);
  const [show, setShow] = useState(false);

  useEffect(() => {
    const delay = setTimeout(() => {
      if (query.length > 1) {
        setLoading(true);
        axios
          .get(`/patient/list?name=${query}`)
          .then(res => {
            setResults(res.data.patients || []);
            setShow(true);
          })
          .catch(() => setResults([]))
          .finally(() => setLoading(false));
      } else {
        setShow(false);
        setResults([]);
      }
    }, 300);

    return () => clearTimeout(delay);
  }, [query]);

  return { results, loading, show, setShow };
}
