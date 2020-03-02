module.exports = file => () => {
    return import (`@/pages/${file}`)
}
