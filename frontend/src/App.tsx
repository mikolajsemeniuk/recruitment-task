import { useState } from 'react';

export default function App() {
  const [index, setIndex] = useState(0);
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  const onSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const data = new FormData(e.currentTarget);
    const number = data.get("number") as string;

    if (!number || !/^\d+$/.test(number)) {
      setError("Please enter a non-negative number");
      return;
    }

    setError("");
    findIndex(number);
  };

  const findIndex = (number: string) => {
    setLoading(true)

    // Add some timeout to simulate the delay.
    setTimeout(() => fetch(`http://localhost:8080/index/${number}`, {
      method: 'GET',
      headers: { Accept: 'application/json, text/plain', },
    })
      .then((response: Response) => {
        if (response.ok)
          return response.json();
        return response.text().then(text => setError(text));
      })
      .then((data: { index: number }) => setIndex(data.index))
      .catch((error) => setError(`An error occurred: ${error}`))
      .finally(() => setLoading(false)), 200)
  };

  return (
    <>
      <div className="bg-gray-100 ">
        <div className="flex min-h-screen items-center justify-center">
          <form onSubmit={onSubmit} className="min-h-1/2 bg-gray-900 border border-gray-900 rounded-2xl">
            <div className="mx-4 sm:mx-24 md:mx-24 lg:mx-36 flex items-center space-y-4 py-16 font-semibold text-gray-500 flex-col">
              <h1 className="text-white text-2xl">
                Find index of your number
              </h1>
              <input className="w-full p-2 bg-gray-900 rounded-md  border border-gray-700 focus:border-blue-700" type="number" placeholder="place your number..." name="number" defaultValue="300" />
              <input className="w-full p-2 bg-gray-900 rounded-md border border-gray-700 " type="text" readOnly placeholder="your number show up here..." value={index} />
              <p className="w-full p-2 text-red-500">
                {error}
              </p>
              <button
                type="submit"
                className="hover:bg-gray-300 cursor-pointer w-full p-2 bg-gray-50 rounded-full font-bold text-gray-800 border border-gray-700"
                disabled={loading}
              >
                {loading ? (
                  <div className="flex items-center justify-center">
                    <svg
                      className="animate-spin h-5 w-5 text-gray-800"
                      xmlns="http://www.w3.org/2000/svg"
                      fill="none"
                      viewBox="0 0 24 24"
                    >
                      <circle
                        className="opacity-25"
                        cx="12"
                        cy="12"
                        r="10"
                        stroke="currentColor"
                        strokeWidth="4"
                      ></circle>
                      <path
                        className="opacity-75"
                        fill="currentColor"
                        d="M4 12a8 8 0 018-8v8z"
                      ></path>
                    </svg>
                    <span className="ml-2">Loading...</span>
                  </div>
                ) : (
                  "Submit"
                )}
              </button>
              <p>
                <span className="text-gray-50">
                  Need to check APIs? Follow:&nbsp;
                </span>
                <a className="font-semibold text-sky-600" href="">
                  http://localhost:8080/#
                </a>
              </p>
            </div>
          </form>
        </div>
      </div>
    </>
  );
}