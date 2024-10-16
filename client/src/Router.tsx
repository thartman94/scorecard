import { RouterProvider, createBrowserRouter } from "react-router-dom";

import { HomePage } from "./pages/Home.tsx";

const router = createBrowserRouter([
    {
        path: "/",
        element: <HomePage />,
    },
]);

export function Router() {
    return <RouterProvider router={router} />;
}
