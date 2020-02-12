package service

import (
	"fmt"
	configs "github.com/fwchen/jellyfish/config"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestImageStorageService_SaveBase64Image(t *testing.T) {
	code := `iVBORw0KGgoAAAANSUhEUgAAAMcAAAEMCAYAAABwXaIFAAAACXBIWXMAAAsTAAALEwEAmpwYAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAACKVSURBVHgB7Z19cBzlfce/z97pJPlVAvNmE/vsYDBtwOLVUJxauJ1CIMWGzDSUEixokk5nMtiQaZo/Wix7pjNp2gHDZPqSpLHspiQk07EJwSSZBsnBTggmWA55ASdBZwLGxm/yq6TT3T59fru3p9XLSXu3z+4+u/d8POd7lW61z3739/bs72HQRE5H54kWYLBl/DtN/V2drf3QRAKDJlDsA384CxhtJniWGcYCcCaeQ9w4CaLFw68RAmE5657xHDfNAwZYL7iZ69pwYS80gaDFIZmOzveyJmtaxZjZJkTQLgSQReAIoTDey0z+LJDqFdYmB41vtDh8YluGgrAKWM2YsQbeLEHAOGIxt3R1XtADTU1ocdRIR+eRdm4IMXCshhKCqEhOjHKPFkr1aHFUgbASWdPAGsbNdVBbEJXIcWZ0GSa2aNdrarQ4PGBZCZZaL+KHdiQFhi5mGhu0SCqjxTEJiRTFOFgP48UN2uUajxbHBNSHKMagLck4tDhcUEzBGd9cV6IYixZJGS0O2OlYEWivjXGgLZs+DmPD1vWtW1DH1L04bBfKENaCKtaaMfQxbqysVytioI5Zs/HEE0IY3dDCqMRCzsy3PrH+2DrUIXVpOezYorhN/Plt0HijDmORuhPH/Z1HV6cYIzdKxxbVU1duVl25VQ9sPLFeCENYDC2MGqkrN6tuLIeILzaDmx3QyICLWG3D1sdaNyDBJF4clKYV8UW3ji8CgOGJLY+d/ygSSqLFoQPvEGDYJgL1h5J4xWJixWELw9Rp2lBgexlnK5MmkESKQwsjCpInkMRlq+wYQwsjfPg13DAfR4JInDjs4FsLIxI4Hlyz8djXkBASJQ6aDqKD74jh6KB6EhJAYmIOGhDGzU5oVIAbhvHg5n+M96zeRIijNCVkGzQqcYKZxZVx7qsVe3HozJTS0Fysa+OawYp9zKGFoTQL45zBirU4SoFfFhp1EQF6XCcqxtatKrlTfdDEgRMl9yqHGBFby1FypzTxoJUzHrv6RyzFod2pOMLbO0RWETEidm6VdqdiDblXi+KSvYqd5TBR7IQmrrSaKMQmOI+V5Si10dGxRryJTXAeK8tht+jUxJxWUft4DDEgNpbjE51HOwy7a4gm/vBS7JGDwsTGchj2qkmaZMDiYD1iYTl0rJFIyHqcp3LmKhaWwwTrgCZpMNUzV8pbDl3XSDRK1z2Utxy6rpFoWop8uAOKorw4GGMroEkqzDDSd0FRlBbH/fZcnCw0CYbmXL2XhYIoLQ6xcbGaqKapCRGYpzugIEqLQ7hUq6BJPGKclXStlBUH1TaglwqQxrQmhjmzVR1u1qaia6WsOHRtQx4kjM8/MAt/v2aWsgIp8pRyLrSy4tBZKjk4wph/cQoXtBiqCkTJrJWS4qDCH3SWyjduYTioKxDeRn2OoRBKiqOIYjs0vphIGA6KCkQIo6BUK1clxcGAdmhqZjJhOKgokCLnWhxTwRiWQlMTXoThoJhAmJFKXQ2FUDQg153Sa6EaYTgoJRCulsegnDhEUKaFUQO1CMNBIYFkVQrKlRNHAcUsNFXhRxju3zGtWYErGPhwFoqgnDgMqBWUeWH+RSn868Mt1n3YyBDG2UGOf956Cm8fKiJqijCUGX8FYw6WRYwgQXyeXBLhmtB9mAJJmjBg1X55FoqgXqmUGQsQExxh0EFKWAdrSAJJoDAsRKl8PhRBQcthxmKy4VhhOIQhkKQKg+Amb4UiaLeqBioJwyFIgSRZGITYo7OhCCrWOZS2HFMJwyEIgSRdGBaMZaEIiVuHPEi8CsNBpkDqQhiKoZQ4SrNxlaRaYTjIEIgWRjRoy+GR+26bXrUwHPwIpA6FkYUiaHF45KlvnRYHVwG1UotAtMWIFi0Oj5wTB9kXtoYnEC2M6FFNHEovhxWWQOpcGDkoglLiiMNacUELRFsMddBuVQ0EJRAtDIIrc4JUsUKeQwyQLRAtjBJcHddaxblVsViGl5ApEC0MGw6chCKoJw4eD8vhIEsgWhgOPAdFUE4cHFyZM4dXZAikVhIWfHPBASiCeuLgZi9iSBQCSWJWKm00KLOKl3riYEYOMSVMgSQ1XcvNYW05KpFGKpaWwyEMgSS5jtG14UJlxl85cZQWbo9NxmoighRIsgt8XKkTo6pN3WJtPYggBJL0yjfnUMalIpQUB+fFfUgAMgVSB1NCRKLK3AmFUFIcJlgPEoIMgdTLXClR6emBQigpDhGU9yBB+BFIHU0i7FcpGCeUFIc9O1et4MwvtQiknmbXCp9KKZeKUHZWrthZzyJhVCOQOpt2zlMstR2Koaw4DMX8T1k4AnkjV5j0M/V2PYbIVClnORRoq12ZNRuOnUCCl1tevjQjbk1Ykk1bz4/2m9i1bwg/+OmgJZC6gfPuLZ1zVkIx0lAYkdrbwpixFgll1768datzuMFSW6AgSl8JKDZOOT9UIx8VXSpCaXF0dV7Qk7SslWY0lHgpTRlSDuWvIU9i1kpThrJUXVAU5cVhIL0JMZ+IqKlIbvP6VmVPfsqLgwqCFJhDkzS4AWMDFCYWrXmE9eiCJmnkVA3EHWIhDmE9RFCenMmIGlCFrUfVQNwhNk3dGC8qbYI1VcGZaWyE4sRGHHZaV1uPJMC4uVF1q0HEqh0o4+xBaOJOH2ISQ8ZKHHS2EZkr7V7FFytDFQerQcSukbSue8QaqmvEJi0fO3FQ3YNx/gg0cYMzbig383YyYrkEQVfnnC4dnMcLEYQ/GRd3yiG263OUgnPtXsUDCsJjFyvGVhx0FmImdHCuPsKdMh+Kw6pdY4n1yk5dG87fxDnX864Uxq5pUI0qfsR+2TMDqXVQaJFFjRu+VwijEzEl9uLYvKIH62/4+c5p6fDXxtBUZk7TED537a+/gxijdIOFqeC7nu8AY+vFw+yugxfgq79aBE300EH1L7f0Yk7zED3tE6fgDezmO2Ln/sZSHLz7+TZk2BPiYbv79e1vXSpu86CJlvsuP4A/m39ozKu8C/n8Bnbr3TnEhFiJg3dva0Fj43pwrKv0maf3L8AP3r4YmmhYvegdcXu30tscBhNW5COxyDLGRhz8pefaYaQ2i4fZqT771V8uwq73LoAmXKYQhps+5IdWqm5FlBeHF2sxEVog4VKFMByUtyJKi6MUW2yDB2sxEVog4VCDMNwoa0XUbSS9+3trhTD2okZhEJ/8w7ew/JIj0ASHT2EQC5FpfI3/eEdVnkEYKGc5LDcq00jWoh2S0FmsYJAgjNEwPIGhoY3Ciigx1UQpcfh1oyZDC0QeVHC974q3g7LKyrhZyoiD/2THGhSxSWxRYF3VXzvSasUh5wpK989WGqp8P7x0P+bPPIcA6YNZfIh9+M97ECFKiIPvfkFUuXknQuDoQCO+8LMrcXSwEZqpKQ6fRbFwFpnmC7Gk9RQ++QdvOZXvoIk8mxW5OIQwRO2CdyBEzhVSlpuli4WTc/bY63i390vW4892fBYfu+IEQiZSgUQmDjvwznSLTWhDRNB8LIpDtBUZDVmLI/ufwbG+58qv/eWK6/Ho3X+KaOCb2S13PoSQicT55t0vZJHhQhjyA+9qWD73iOUqkECcekh+4H3LhahXHGsxLPaDw8zmRlxy3mxEB3uQ7xbJmnx+ZZiZrNAthyrCGAtZkW//sgE/fv5htFy6Ehdc8fG6EsnZY7+wrAXdu7nusvlYf9+dEYvDge8NUyChikNVYTis2vjvOHj8ZPl5PYiELOWRN59B/zsvjnqdrAW5UR+98SqohSWQe8JI9YYmDtWFcXpgEI9v+yG++8rr495LokjIQhx76zmcPvzKuPfuFfHFp29fLgTSBDUJx4KEIg7VheFm/7uHLZH87Ldvj3uPRNLygVsx/fwPIa5Ucp8ItVyoqQheIIGLI07CcEMW5Bs7X7XEMpYGYUHOX/hRzLxkWSysCQnh3LFf4qiwFKaoWYyFRPEpYSnoPl4EK5BAxaFCutYvZEG+8r1dE1oSgqzIzItuVE4ojiDofiIrQbRfdbnlQsVPFG6CS/MGK44ICnxB8Z4I1L8sRPKaEIk7aHdDFsUSy8U3omn2wlDFMnzufZw5LgRx9Bc4deiVCS0EQYH2vStusOoW6sYUVcLwBPujOx6FZKSLY/D+ZWuLRX5g2mc2LA1rSkjYkMtFt0rWxIHE0ixE0jgri6ZZQizTLrReSzVMR61QgY5qEAOn+jB0Mmc9PiMEUUkMBAniussWJMBKVISLI/lRIZBNkIhUcZy574aOFDPoUlakb2gXtxVIMmRNSCA7X99v3Z8e8DbnyEhPt4RCIqGbkZ5Wfq9BvE5WwMEsnLMEQSnXYv7spCJwQ4JYIdwmEsOKqxYnx0pUhsMsrpQ5WVGaOAY6bsrygrmXcWbNqk0tvAINH7kX9QQJhG7kelEg71UsMrCtw3xcK26Xz7soqRZiKk4gP3StrBqIFHHwjvaWweGB8lV7bKaIw1evse7rGRLIweOnhGAO4Dfvvm+JZaLsVzWQCMgKkAhmiMdXlIQQj/RrKPSVBOI7gyVFHIN/tewJDrsBAmtsQuYv/qbuhTEZVHAkl4zEQrcz4rkDBftzXQf6DCEEEgS9NqMkDM0USArQfYvDHWcQ6eW3IX31TdBoIkRKgO5LHBRnYHikwJe6Yika/mQ1fNHYbN8PDUCj8YHv+MNf95FhXm6yRm5Uevnt8EU6DfzhMuBDy0ZEotHURisyjV+DD2oWB7lTcHUIsQLwRp/+8AevBpqabWGQQDQaf7T7aflTkzjInRJxxnrnOdU0fAfgH1gMnH/RyPODOWg0PmEi+niMd2/LogZqEoc5bHbC7U75LfZdkhXiuGzk+e9/IypsOWg0EqjZvapaHGf/8sbVBtga53nDylXwBblQC68cef7+O0Icv4VGI5Ga3KuqxCHiDBiGtS6GBWWnjHlZ1IwTgDtQhir3a2g0knHcq6p8f8/iEHEG3XXA7U7d2A5fZP/ADsCJQgH4xU/te41GPq1obHysmh+oxnKMCsJTS5f5C8IvnGffHHK/0rUNTbBwrOO7nvd8bZEncZDVMApjrIafKjjFGZcuHnn+Xk7EGhIbEms0E8PA2ONeP+zVcmRNzl2pW5/ZKUrbNrkq4ZSd0mjCod1aJcwDU4qDrEYpdWtBViO1xMdVr2PdqV//TMcZmjBhMFKeYg8vliPrTt36shoTuVPnTkOjCRlP1mNKcbAiX1t+7NtqXDrandJVcE00eLIek4qD6hrcRHmarW+rMbYKrrNTmuiY0npMZTk64MpQ+bIaH3C5U2dP6eyUJmoYUqmHJ/tARXEM3H8TRk8u9Gk13EH4G69Bo4kcjtWTVc0nsxztcFkNY+ES1IzbatDcKe1OadSAIZOpOOdqQnFY6VtudpQ/NHdB7ddqjLUaelKhRilYRdeqkuUYnb71M4dKWw2N2rRUCswnFEcxXyx/2JibrX0OVR1YjaH/fhKFXd8HHxqE6vCjh5DfvgWFn78MTZmKad1x4iCXSgTiZauRWrIUNTP7vJHq97HDibMaRXGQ8dP91sGW/9Z/ovhGL1SEhFvYsxNDYhtNUVuix/xUaKuHxYH2iQLzidYEzMJ1bbiv6zUoXXvyuC2SM6eQNEzXAUYiGX7xWZjvHrAye2yWGn27zHdz1nbR9jlQyxkSSWpWbJvfy6ehsUP8P6qVzzhxkEslLIf12JdL5UDWIqE1jYblt8OYc7F9Ji4dfMU3e+0Db+mySPt3OdaiOMaFouQKtU/STfdGwcS/uzBGHKP6Vp3oaEPTcCP1oWqn53QJrK/CX51guVZ7eoRbtW/U61Zb1FUPCCvSijAhQViCdcVBlG1MCYumG+5VhCM/dJ67jejYmINOJ+3lN/24VHUEiaBh5epxbVBJNENff8o6UMOAvo8C7uExCQJq6p25f60WxlSkMqM6Eo5yqzKFTLvzWIpLVWeQi9X4ibWWFXELwrYqvcINu81fMXUSHBdqlLWwRLtKn+S8wZDCH4v7LueF0TEHH5lkaMxbAE1tUB8vckfzO54BP3bIes06q7/wDFJXtEkN2K2AmyxF6XscUsJKWN/TqBtPe4fR8V9eQq0cc1gNFIZ5H0pTRjKr1ugzjgTIYrgDdkLWEg2FfSK22P39Ua9ZAfeHbwc7/2JoqoaDp69ly//Mysm7Y44snBVfxdlGC0MOZEFICKPqRZkmKS5ratGSsmVgouBKHe4zqzu0MHxRaHceld0qs2C2GSVDYuidKxUnYKc4jqxI5o6PQwbWZQRXLxOp4wNWbKFjRN9QAwaafm6ldMviYCZrd5wsHW8EA1kR2alxim80MuEr+E+2gd1894hbxRjKdp/OcBpNnULmN0sPLHHk7/sjikTKpzQ2R7tVmjrGbLYMhWM52oSzZTusIsDT6T9NHcPAi7fSA0scw0Yh67yjg3FN3cMMK+i2A3ITbU4wrl0q73AOvHfcrl/QUseMSVvWXQpcbCCtWksB5SWttH3QeIJbbpUljhTYUl56mc3S61lPxamBQTyz81V8Y+cea6lkh4/eeBU+ddstmHt+tCnVU+cG8OXv7caOPa+Xt4+Wa15x1eX41O3LRy3lrJmQrJWxokcDf7WsPBNXV8Ynh9YJ/9svPW3dV2L9fXdaQomCN985hM99bVvF7SML9x+fuU8LZHI4jKFFTkCedV5lTQoE43T14NkzwLGjwOFD9o0en+yPvK/u3/3X/04qDOLxbf+H/e8eRtgcPNY/qTCI90riPj0Q8WW9Co+xRSGdTdM1HBh2iSOqgJx2ivs2GZRNmy1cl/Pn2LeQ+O4rr4uD/v0pP0euzOPbfmidocPkK9/fPaVwCfrMN4Rb+GnhYoVKLWNMt4siOCa5kaWYo2XUBoUN7aS3c1PvLDc0Lfv9Q/aNtpl2Ht0C3v6e1/d7/uzPfvu2dXae2RzOPqXkgBfhOuwUf0to4vA7xvSzIY1xGYO1pKfnp2WLrGg/D3NuzqD443/zRnU7bCJoJ9LOI7M8PxvoWeaMK/j2+vmwxEFucjWu3G8OehdSzcRwjEuIWJwvSBcNswVOqioT0kAefMf+Y2X6lkOlgTgu/NZFl0VjBWMEZa8CJcgxJmty+ZJgx5hhtmHCDDfvSDvsrd8GF3RRUPfzXntHSmbxvAs9f5YOvktCzQgxXHfZfM+fvnxegGffoMeYLFFAYzyC0Wpwk5fFwWYGPJi0w2jHBQ3ttAB2XruoE3hlRRWflQEV+Lymjyl//9EbPoRAiPkYj8BnV7UOuS9oh5GpDYsAdh6dme9dcf2Un6MaQuiZIMGdN1zlyXpct3gB7gyiDpOAMXYTjjgOHwrnbDIW2mn734BMPnv3n1pV5kozMeaWimyXRFBkI+vxxb++p6JA6P3rhTC++NDdkE6Cxtgh9Q9Xf6Bd7LPbQNOt5i1Eav5lkIqTsShGVNihnZcWGeuZsyALOvhWXLUY+ULR6gbW1JDG9eK1e265Bv+0ZlWIGarxNIptIfeKRMqs5w1oyjSg/erFeGS1LWz6jFQSOMYEG7j3pizSyGbueqCbOmJIv9Ryfym7ECW0466/yb7XyCeZY9xX9g747h3lziPSoDPKq4p09Kb8ON00cknuGHcHG3NE4YNWggJFvd65fBI8xsGKw29lVCa006I2/UkkwWMcnDhop6m2oAsVCDXySPgYj4iDIweZqHgg0hRpjTwSPcY8NyIOBrn2UcUDkcxuDJYniw2qjrGM7eI4mXY/gcxrjAcVPQj7+8O9PqBYFPtiwL45wsznIZ1MRvgBKUDUNdDUDEyfgcBRdYzPnPH/93P0u8TBc1KvwK/3MzSdvY4ctkVhFhE458Y8b8jYB8ici2zhBIGqY+y/GMnFvwNuy5FDPRD0gDqiOBexyzEsrFP/cfvWcl6wIlENGelcZuZc5cRi/8TrZ2o8c+ggcPwIlIMEQqK9eC4wUzdW8ARr7B8JyBsaeiETVadqBHGBDMUVud+pKQwHsiS/z9lWTRYJHmNao2OkkfTNH8lBJilFd5zs7iokjAO/i96N8sqRQ/IEktwxztF/xkQvSmFGCNmSWpB9tnOC7jhBAiFXyy/JHeMc/TdWHPJcq9nRdv2bEDK3MlOcdICp7EpNxuGD/lPKSR1jzq01s0eLg0PemsDTFTyryBxMOrCOhN+4TRrkDh78PXyRzDHmYKkeejBaHGYhB1nQRqrWAURmAzgSxnAAxbwwoTjJTzWZxlg1gcgYY24coLvR4igWeyCTKDrVVYKEKlMccQnAp+K4z/lRIXacnBI5Y9w/0WqyYLfeTfOrcpDF3EvVSffJvNCJYo24Ww0HEnnRRwU/eWNcjrvHT1mXGXfQTqOdFzVOy1BZnD2LxEDC8ONa0RircIWlrDE2efn4NyZ4swcyoR0XtV9K3fFkErfU7VT4dRHpBBj1GMsRKBcCmEQcxfx2yObKD0VnemmnyU45DiVMHDJcxKjHWI5n0M8+/Oc9zpNxfw3FHXz388LvYm2QBVUsF15mt28JExKFZJPPT/cj//IPoAqsaRoybT4byOWH4ZtkjPGoOt/EUjdF3GFAnjgIR9lh7Twy81cG0/KSD55DopA1pT7eY8zFMd/lfqHCNeSmfNeKoJ23eEnw5vdC8T1Xtek+VV6hi6RkQWMchotFKVvpYzz0I/ezilc38d07TsC9sI1M6Aqy1wPqcUrLDwSZIRPZHb7vFagEuVa+mDYdyAbQ6TJeY9zNbrljpfuFyrIz+RYYbC2CgPzTG24aaTws4+IU8j2DXrOBSKXsrpBJqXMQdFmt9N/pGmPqoytDJDTGJAz5mbFxLhUxiU0i1yoVjDgcKJAiF+j9Q7XvQCcgC3MSHB1MWhzeiM0Yj3apiEkvGg+kRehkUB8kavdytjTnZyKL4iykSNOlaadHEVccO2LPak0Ki6+0rzkPAzXHeJxLRUz+rbZrtR5h4awe6sZ9plFlIiNdk330sL9pF6pA8UZDiNeWqzfGE7pUxOQdDwv5TYga2lnOTRVE3CG73X1kkNCjJtoxzrGb79g60RuTisOaiMjRA814LlBoxnGtkMVQQRyRUnm61NS9cnlxAzTjoQPrgosQa7IfRJ0jXKr8xkpvTikOa66Jth4TQ9ajMcBMT5CQsBvqpI9VRdiz7Oa7c5Xe9dhlXQTmmomhFGPcDrLz5iTDLfQHh1l4crIPeO7/GWjFPO5QzYP6VsWh9jG7FZjnfb3yBLNXpG+vnewD3tfnMPmT0EwMWQ7y36cp2qrGgVwpLQyC0rdTZmK9W47ubS3INL4lHrZCUxm6hFa15gtUyyA3SsVuIdHQJ6zGoqk+5Nly2Gld/ig0k0OpUbIiF82NPhYhUSz4oD2pUAvDgaNYeMTLB6tec0DEHmQ9FkLjDZoiQZeh0qW11OsqyOUIaOp5c7OdQWs9T+5U9JjDT59AYd8rSF95Tc/AN//t1ulPTz2zuvpJK2bxIbHTu6HxBp2x9Vk7copv9KL485fp1i6e0pSoKet3VS+YWap7aHFoYgM/JazGnvKkW97AjD4vP1fjarI69tDEBM5ReHVkNrqII/Zm/uflrV5+tCZxsOV39orUrp5WolGe4pu9lktVgvMi+5jXn619HXJ7xq4n86TRRMFYd0oYkSebv/lyzuvP1ywOK7VLwblGoyIld4paKdnPkWs+11yVt1O75UApODd59Nd8aDRjGOtOpRjrZNt7+qv5Hb7EYVHIkxq1e6VRhrHulLh1eQ3C3fgWh3avNEoh3KnhF78zyp1CkW2s5Vf5txwou1c6e6WJnMKr4lA8mHOechjsoWqCcDdSxEGwD9/ZqYuDmigx3+kb5U6ZnG9o/vrLPagRaeKwSDFyr3T8oQkdijOGu7/jfqln+tOv+PJmpIrDWsuc83tgB0EaTSjwwQHkn93qjjP6RJzhOw6WazlQqp5z6OklmnCgeka3KwCnliAw76k1znAjXRwEW37HJh2gawLHKvT1oNj3husVPDLj6T29kEDV13NUA9+1Y7P4hg5oNAFQ2NM9LgD3G2e4CcRylBkeekRs815oNJIp/nqvWxjgjG+XKQwiUHFYBcJ8nhr0VlW212gmhQp97syUCMBZwZAe5wZrOeBU0HXnEo08aN6UAwc/AZOtlBGAjyVwcVjY09u19dBIodi3v/yYgW0PQhhEKOIoNaSWkkHQaPjRQ+WHqbz5FAIiHMthcwAajV9EvOGqaSDz7T2BnXTDFIdGI52Be29CUIQiDt69je4WQKPxC2P2gqUjZBEQYVmOFhE5tUOjkQCbU+4Qz0zD7EBABC4O3v0CtcVcDY1GEqkrlpYfG2CfQECEsRRrVpjC8Bbd1CSe1KIlKJBrlWnixszzPgb8FEEQ7Nyql54T35DqEt+yBhqNRChjZcUeDJswNPSIKBdANkG7Ve1aGJogKAflHGsRUFAerFvFUh2IgOEXvgk+NCR24Gykl98OptIyzRXgQ4Niu5+xn4jtzXzk44gD5rs5qwUOP30SxrwFSN/QjpBhIqZdJ+7XQTKBWQ6+63na7FWIgGLfm9ZF9sU398VCGBb5QWub6eaqACsPuTe0n63tPhXRDCHG7iqVC6QSpFvVhgjWEKQzcJm4CCPGuGsOZD0iIosAjrXgxMHNSBbXdE8tMGbq9T2Dhs2aXX7s3veh0zA9C8kEKA4jiyhwW46MthyB49rHPD+IyOBDMbIcKEZyGhl19mpshCZYmPsENBShOAIgOHE0NEQxRZ0LcVBbIOvGZraWH0Nt+CQ3leEiGBb7uaW8vSIoj2a7CwXpx1tgqVzqYcV376DTeNCOPxdDQjtmJ0zsS1+/ojfNjf6ze37Yks40t1gLVBoNS0UMtMo1vyvQ4qdH6ACiPl/P0nazGbN7UWSW2eP9J2ntxSwYuaZshdjaFbCDTjW2m1sXrvWK7duOYv5tfuakSFmxXNNN7cCMmVkUCy0w0qtC3O5e65ohyQTdfSS46jjHCfHfFnHQP2v16vXyIz95QQycKQqT7DFEtyKuEAXvgmlu9brd1g+99Fy7VTdieADRiMQWM4xNmDawlV0z9cFopfOtrCVbF+x2883sljulNzMPevpIewArz9KZqwepoYfYzXfnUCP8pec7YbCHxcNWhIPv7RbipruFMPkPEaa4rRMRNrDld9TUC6C03eLExDcHMDubwxha5OdYqETgZyBhPV4U33IrZGAP0karaZyMX0eWxOQvIvgDjYTxqLTtfsk6I5O4yQIGOYbcaq1k5D8m4+ALYLstKxyE1SCCF4d9AL4Gv2do241aabUblYiorLYgkxECYdcgCGi7efGealwoT7921w66I3flcQQzjiToLULQD0Ii/EdWS53VSKW/Bv9Wu09YjZVBWA0iFN+15F696OP7At0JRCDdGQMSdPnX2z69iEMYCUSme0hJjifFdj+CACht9zViu8k9rG27OY6L//4kqH1LhNN9hM6anF+L2pYnCFwYhHWGtPv7SkpDCnckNXRtkIMnfjeQMrqEm3IdZC39YFk6PBKUMAhru4G94pighn/Vbje3fyZYYRChZj3sbBFf78pgTfb9tBO2Y5oIYK+Rn6ar+KUjcUgWtewf6+DiT1mL+YREadJdC9KZdT78eSmJjqq+0AnUTf6EuKdJqmyK7esni4YZ+SfDOCYiyZuXU6rAw8K0to1+k3Lo1aVog0CY/g6xex7wUBvhpf9p4J4SA7cpTDGP2pCRrNB6jzUGZ9uFZS9ujGp/WxfFWdf+pDrE/V1ii0fXxqzt4zvDEoVD5EUlvlcExGczWetJKt8f1lnLK5aQh4fbRMwkDja2dNSbjOfERu+DObwvSiFPRPmAgyHiPXa1OMBGDjhmCfmAKD72Ymb+2ajEPBH8J5YVzKKQpiJoP6bnc1Ft3/8DtAf5i52RYCAAAAAASUVORK5CYII=`
	cfg, _ := configs.LoadConfig("../config/config.yaml")
	storageService := NewStorageService(&cfg.Storage)
	err := storageService.init()
	assert.Nil(t, err)
	bucketName := fmt.Sprintf("jellyfish-image-storage-test-%d", time.Now().Nanosecond())
	imageStorageService, err := NewImageStorageService(bucketName, storageService)
	assert.Nil(t, err)
	fileName, err := imageStorageService.SaveBase64Image(code)
	assert.Nil(t, err)
	_ = imageStorageService.storageService.client.RemoveObject(bucketName, fileName)
	_ = imageStorageService.storageService.client.RemoveBucket(bucketName)
}
