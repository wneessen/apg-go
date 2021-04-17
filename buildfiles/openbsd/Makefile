COMMENT =		"automated password generator" clone written in Go

V =			0.3.1

DISTNAME =		v$V

REVISION =		1

PKGNAME =		apg-go-$V

CATEGORIES =		security

HOMEPAGE =		https://github.com/wneessen/apg-go

MAINTAINER =		Winni Neessen <wn@neessen.net>

MASTER_SITES =		https://github.com/wneessen/apg-go/archive/refs/tags/

# MIT
PERMIT_PACKAGE =	Yes

MODULES =		lang/go
MODGO_TYPE =		bin

ALL_TARGET =		wneessen/apg-go

post-extract:
	mv ${WRKDIR}/${PKGNAME} ${WRKDIR}/${DISTNAME}

test:
	cd ${WRKDIR} && ${MODGO_CMD} test ${ALL_TARGET}

build:
	mkdir -p ${MODGO_WORKSPACE}/bin
	cd ${WRKDIR} && ${MODGO_CMD} build \
		-o ${MODGO_WORKSPACE}/bin ${ALL_TARGET}

.include <bsd.port.mk>
