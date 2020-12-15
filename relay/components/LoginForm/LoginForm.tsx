import React, { useEffect, useState } from 'react';

import { Magic } from 'magic-sdk';

function LoginForm({ magic, setLoggedIn }: {magic: Magic, setLoggedIn: any}) {

    const [email, setEmail] = useState("");
    const [loadedMagic, setLoadedMagic] = useState(false);

    useEffect(() => {
        (async function() {
            if (magic) {
                setLoadedMagic(true);
            }
        })();
    }, [magic])

    const handleLogin = async () => {
        await magic.auth.loginWithMagicLink({ email });
        setLoggedIn(await magic.user.isLoggedIn());
    }

    return loadedMagic ? (
        <div>
            <h1>Please sign up or login</h1>
            <input onChange={(e) => setEmail(e.target.value)} type="email" name="email" required placeholder="Enter your email" />
            <button onClick={handleLogin} type="submit">Login</button>
        </div>
    ) : <div>Loading...</div>
}

export default LoginForm;
