"use client"
import React, {useEffect, useState} from 'react';
import {useLogin, useLogout, usePrivy, useIdentityToken} from "@privy-io/react-auth";


function PrivyReadyComponent() {

    const {logout} = useLogout();
    const {login} = useLogin();
    const identityToken = useIdentityToken();
    let [isNewUser, setIsNewUser] = useState<boolean | null>(null); // null znači "još ne znamo"
    let [fetched, setIsFetched] = useState(false);
    const {ready, authenticated, user} = usePrivy();


    useEffect(() => {
        if (identityToken !== null && identityToken.identityToken !== null && !fetched) {
            const fetchProtectedData = async () => {
                const token = identityToken!;
                if (!token) {
                    console.log("ERROR");
                    return;
                }
                setIsFetched(true);
                const response = await fetch(`${process.env.NEXT_PUBLIC_SERVER_URL}/api/auth`, {
                    method: "GET",
                    headers: {
                        "Authorization": `Bearer ${token.identityToken}`,
                        "Content-Type": "application/json"
                    }
                });

                setIsNewUser(response.status === 201);
                const data = await response.json();
                console.log("Protected response:", data);
            };

            fetchProtectedData();
        }
    }, [identityToken]);


    if(!ready)
        return (<div>Getting Privy Ready</div>);

    if(!authenticated)
        return (<div className="mt-6 flex justify-center text-center">
            <button
                className="bg-violet-600 hover:bg-violet-700 py-3 px-6 text-white rounded-lg"
                onClick={login}>
                Log in
            </button>
        </div>);

    return (
        <>
            <div className={"mt-6 flex-row justify-center text-center"}>
            <div>New user: {(isNewUser === true? "New" :"No new")}</div>
            <div>Successfully authenticated {user?.wallet?.address}</div>
            <button
                className="bg-violet-600 hover:bg-violet-700 py-3 px-6 text-white rounded-lg mt-4"
                onClick={logout}>
                Log out
            </button>
            </div>
        </>
    );
}

export default PrivyReadyComponent;