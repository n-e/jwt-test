// openssl genpkey -algorithm Ed25519 -out k.pem
// openssl pkey -in k.pem -pubout
// spec: https://datatracker.ietf.org/doc/html/rfc8037
// rationale: https://datatracker.ietf.org/doc/html/rfc8032

var jose = require("jose")

const pub = `-----BEGIN PUBLIC KEY-----
MCowBQYDK2VwAyEA6FZI/NownlUV6kQ4PoqZG1ZKtGK7ARdRNFdN9azzEmw=
-----END PUBLIC KEY-----
`

const priv = `-----BEGIN PRIVATE KEY-----
MC4CAQAwBQYDK2VwBCIEIA4ldGTDlpZSSVrlygitLuEnpKz3zAFcCfolh75DfE1b
-----END PRIVATE KEY-----
`;

(async () => {
    const pubx = await jose.importSPKI(pub)
    const privx = await jose.importPKCS8(priv)

    const sgn = await new jose.SignJWT({ a: 'b' })
        .setProtectedHeader({ alg: 'EdDSA' })
        // .setIssuedAt()
        .setExpirationTime(Date.now() / 1000 + 100000)
        .sign(privx)

    console.log(sgn)

    const ver = await jose.jwtVerify(sgn, pubx, {
        algorithms: ['EdDSA']
    })

    console.log(ver)
})().catch(e => console.log(e))